package main

import (
	"crypto/rand"
	mrand "math/rand"
	"sort"
	"sync"
	"time"
)

// Phase represents the lifecycle stage of a quiz.
const (
	PhaseLobby    = "lobby"    // waiting for teams to join
	PhaseQuestion = "question" // a question is live, teams can answer
	PhaseRevealed = "revealed" // the answer is shown, scores updated
	PhaseFinished = "finished" // all questions done, final leaderboard
)

// Team is a group of attendees playing together, run by a captain.
type Team struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Score    int       `json:"score"`
	JoinedAt time.Time `json:"-"`
	// Difficulty is the team's own difficulty in a mixed-mode quiz.
	Difficulty string `json:"-"`
	// answers maps question index -> chosen option index for this team.
	answers map[int]int
	// questions is this team's own question stream (mixed mode only).
	questions []Question
}

// Quiz is a single live quiz session held entirely in memory.
type Quiz struct {
	ID         string
	AdminToken string
	Difficulty string
	Lang       string // language for question content: en, is, sv
	Source     string // question source: official | community
	Mixed      bool   // mixed mode: each team plays its own difficulty
	Rounds     int    // number of questions per game (used in mixed mode)
	Questions  []Question
	Points     int

	Phase   string
	Current int // index of the current question (valid once started)

	teams map[string]*Team
	order []string // team ids in join order

	CreatedAt time.Time
	mu        sync.Mutex
}

// mixedRounds is how many questions a mixed-mode game plays. It matches the
// smallest difficulty draw so every team's stream is the same length.
const mixedRounds = 12

// mixedPoints is awarded for every correct answer in mixed mode, regardless of
// difficulty, so a Kids team competes on equal terms with Adult and Nerd teams.
const mixedPoints = 100

// DiffMixed is the create-time value selecting mixed per-team difficulty.
const DiffMixed = "mixed"

// Question source: the built-in bank or the community-voted board. Chosen by the
// quizmaster in the lobby, like the language.
const (
	SourceOfficial  = "official"
	SourceCommunity = "community"
)

// normalizeSource coerces arbitrary input to a supported source, defaulting to
// the built-in bank.
func normalizeSource(s string) string {
	if s == SourceCommunity {
		return SourceCommunity
	}
	return SourceOfficial
}

// Store holds all quizzes in memory. Everything is wiped on restart.
type Store struct {
	mu      sync.RWMutex
	quizzes map[string]*Quiz
}

// NewStore creates an empty in-memory store.
func NewStore() *Store {
	return &Store{quizzes: make(map[string]*Quiz)}
}

const (
	idAlphabet  = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789" // no ambiguous chars (0/O, 1/I)
	hexAlphabet = "0123456789abcdef"
)

func randString(alphabet string, n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		// crypto/rand should never fail; fall back to time-based jitter.
		for i := range b {
			b[i] = byte(time.Now().UnixNano() >> (i % 8))
		}
	}
	for i := range b {
		b[i] = alphabet[int(b[i])%len(alphabet)]
	}
	return string(b)
}

// CreateQuiz builds a new quiz for the given difficulty and language and
// returns it. A difficulty of DiffMixed creates a mixed-mode quiz where each
// team is assigned its own difficulty in the lobby.
func (s *Store) CreateQuiz(difficulty, lang string) *Quiz {
	lang = normalizeLang(lang)

	s.mu.Lock()
	defer s.mu.Unlock()

	var id string
	for {
		id = randString(idAlphabet, 4)
		if _, exists := s.quizzes[id]; !exists {
			break
		}
	}

	q := &Quiz{
		ID:         id,
		AdminToken: randString(hexAlphabet, 24),
		Lang:       lang,
		Source:     SourceOfficial,
		Phase:      PhaseLobby,
		Current:    0,
		teams:      make(map[string]*Team),
		CreatedAt:  time.Now(),
	}

	if difficulty == DiffMixed {
		q.Mixed = true
		q.Difficulty = DiffMixed
		q.Rounds = mixedRounds
		q.Points = mixedPoints
	} else {
		meta, ok := difficulties[difficulty]
		if !ok {
			meta = difficulties[DiffAdult]
		}
		q.Difficulty = meta.Key
		q.Questions = pickQuestions(meta)
		q.Points = meta.Points
	}

	s.quizzes[id] = q
	return q
}

// pickQuestions returns a freshly shuffled subset of the difficulty's pool, so
// each game is different. If PlayCount is 0 or exceeds the pool, all questions
// are used.
func pickQuestions(meta difficultyMeta) []Question {
	return pickN(meta, meta.PlayCount)
}

// pickN returns a freshly shuffled subset of n questions from the pool (capped
// at the pool size). n <= 0 means the whole pool.
func pickN(meta difficultyMeta, n int) []Question {
	pool := make([]Question, len(meta.Questions))
	copy(pool, meta.Questions)
	mrand.Shuffle(len(pool), func(i, j int) { pool[i], pool[j] = pool[j], pool[i] })
	if n <= 0 || n > len(pool) {
		n = len(pool)
	}
	return pool[:n]
}

// Get returns the quiz with the given id, if any.
func (s *Store) Get(id string) (*Quiz, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	q, ok := s.quizzes[id]
	return q, ok
}

// AddTeam registers a new team (captain) and returns it.
func (q *Quiz) AddTeam(name string) *Team {
	q.mu.Lock()
	defer q.mu.Unlock()

	t := &Team{
		ID:       randString(hexAlphabet, 16),
		Name:     name,
		JoinedAt: time.Now(),
		answers:  make(map[int]int),
	}
	if q.Mixed {
		// Default new teams to Kids until the quizmaster assigns otherwise, so
		// a team always has a playable question stream.
		t.Difficulty = DiffKids
		t.questions = q.drawForLocked(DiffKids, q.Rounds)
	}
	q.teams[t.ID] = t
	q.order = append(q.order, t.ID)
	return t
}

// SetLang changes the quiz's question-content language. Built-in questions are
// localized on render, so the change takes effect on the next poll for every
// device. Community questions are language-specific and drawn ahead of time, so
// in the lobby we re-draw them to match the newly selected language.
func (q *Quiz) SetLang(lang string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Lang = normalizeLang(lang)
	if q.Source == SourceCommunity && q.Phase == PhaseLobby {
		q.redrawLocked()
	}
}

// SetSource switches between the built-in bank and the community board as the
// question source. Only allowed in the lobby (questions are dealt at start), and
// it re-draws the question set so the change is immediate. Caller-facing.
func (q *Quiz) SetSource(source string) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.Phase != PhaseLobby {
		return false
	}
	source = normalizeSource(source)
	if source == q.Source {
		return true
	}
	q.Source = source
	q.redrawLocked()
	return true
}

// drawForLocked draws n questions for the given difficulty using the quiz's
// current source and language. Caller must hold q.mu.
func (q *Quiz) drawForLocked(difficulty string, n int) []Question {
	return drawQuestions(q.Source, q.Lang, difficulty, n)
}

// redrawLocked re-deals the quiz's questions for the current source/language.
// For mixed mode every team's stream is redrawn for its own difficulty. Caller
// must hold q.mu and should only call this in the lobby.
func (q *Quiz) redrawLocked() {
	if q.Mixed {
		for _, t := range q.teams {
			if t.Difficulty == "" {
				t.Difficulty = DiffKids
			}
			t.questions = q.drawForLocked(t.Difficulty, q.Rounds)
		}
		return
	}
	q.Questions = q.drawForLocked(q.Difficulty, difficulties[q.Difficulty].PlayCount)
}

// SetTeamDifficulty assigns a team's difficulty in a mixed-mode quiz and
// (re)draws its question stream. Only allowed in the lobby.
func (q *Quiz) SetTeamDifficulty(teamID, difficulty string) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	if !q.Mixed || q.Phase != PhaseLobby {
		return false
	}
	meta, ok := difficulties[difficulty]
	if !ok {
		return false
	}
	t, ok := q.teams[teamID]
	if !ok {
		return false
	}
	t.Difficulty = meta.Key
	t.questions = q.drawForLocked(meta.Key, q.Rounds)
	return true
}

// questionForLocked returns the current question for a team. In mixed mode that
// is the team's own stream; otherwise it is the shared quiz question. Caller
// must hold q.mu.
func (q *Quiz) questionForLocked(t *Team) (Question, bool) {
	if q.Mixed {
		if t != nil && q.Current >= 0 && q.Current < len(t.questions) {
			return t.questions[q.Current], true
		}
		return Question{}, false
	}
	if q.Current >= 0 && q.Current < len(q.Questions) {
		return q.Questions[q.Current], true
	}
	return Question{}, false
}

// totalLocked returns the number of questions in the game. Caller must hold q.mu.
func (q *Quiz) totalLocked() int {
	if q.Mixed {
		return q.Rounds
	}
	return len(q.Questions)
}

// Team returns a team by id.
func (q *Quiz) Team(id string) (*Team, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	t, ok := q.teams[id]
	return t, ok
}

// SubmitAnswer records a team's answer for the current question. It only
// applies while the question phase is live and the team hasn't locked in yet.
func (q *Quiz) SubmitAnswer(teamID string, choice int) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.Phase != PhaseQuestion {
		return false
	}
	t, ok := q.teams[teamID]
	if !ok {
		return false
	}
	que, ok := q.questionForLocked(t)
	if !ok {
		return false
	}
	if choice < 0 || choice >= len(que.Options) {
		return false
	}
	if _, answered := t.answers[q.Current]; answered {
		return false // already locked in
	}
	t.answers[q.Current] = choice
	return true
}

// teamView is a leaderboard entry safe to expose to clients.
type teamView struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Score      int    `json:"score"`
	Answered   bool   `json:"answered"`             // answered the current question
	Choice     int    `json:"choice"`               // -1 if not answered / not revealed
	Correct    bool   `json:"correct"`              // whether their answer was correct (revealed only)
	Difficulty string `json:"difficulty,omitempty"` // assigned difficulty (mixed mode)
}

// leaderboard returns teams sorted by score (desc), then join order. Caller
// must hold q.mu.
func (q *Quiz) leaderboardLocked(includeChoice, revealed bool) []teamView {
	views := make([]teamView, 0, len(q.order))
	for _, id := range q.order {
		t := q.teams[id]
		choice := -1
		correct := false
		_, answered := t.answers[q.Current]
		if includeChoice && revealed {
			if c, ok := t.answers[q.Current]; ok {
				choice = c
				if que, ok := q.questionForLocked(t); ok {
					correct = c == que.Answer
				}
			}
		}
		views = append(views, teamView{
			ID:         t.ID,
			Name:       t.Name,
			Score:      t.Score,
			Answered:   answered,
			Choice:     choice,
			Correct:    correct,
			Difficulty: t.Difficulty,
		})
	}
	sort.SliceStable(views, func(i, j int) bool {
		return views[i].Score > views[j].Score
	})
	return views
}

// Start moves the quiz from the lobby into the first question.
func (q *Quiz) Start() {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.Phase == PhaseLobby && len(q.teams) > 0 {
		if q.Mixed {
			// Make sure every team has a question stream before kickoff.
			for _, t := range q.teams {
				if len(t.questions) == 0 {
					if t.Difficulty == "" {
						t.Difficulty = DiffKids
					}
					t.questions = q.drawForLocked(t.Difficulty, q.Rounds)
				}
			}
		}
		q.Phase = PhaseQuestion
		q.Current = 0
	}
}

// Reveal scores the current question and shows the answer. Correct teams each
// earn q.Points — equal for everyone, even when difficulties differ.
func (q *Quiz) Reveal() {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.Phase != PhaseQuestion {
		return
	}
	for _, t := range q.teams {
		que, ok := q.questionForLocked(t)
		if !ok {
			continue
		}
		if c, ok := t.answers[q.Current]; ok && c == que.Answer {
			t.Score += q.Points
		}
	}
	q.Phase = PhaseRevealed
}

// Next advances to the next question, or finishes the quiz.
func (q *Quiz) Next() {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.Phase != PhaseRevealed {
		return
	}
	if q.Current+1 >= q.totalLocked() {
		q.Phase = PhaseFinished
		return
	}
	q.Current++
	q.Phase = PhaseQuestion
}

// Adjust nudges a team's score by delta (admin manual override).
func (q *Quiz) Adjust(teamID string, delta int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if t, ok := q.teams[teamID]; ok {
		t.Score += delta
	}
}

// Reset returns a finished/in-progress quiz back to the lobby, keeping teams
// but clearing scores and answers.
func (q *Quiz) Reset() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Phase = PhaseLobby
	q.Current = 0
	for _, t := range q.teams {
		t.Score = 0
		t.answers = make(map[int]int)
		if q.Mixed {
			// Keep each team's assigned difficulty but draw a fresh stream.
			if t.Difficulty == "" {
				t.Difficulty = DiffKids
			}
			t.questions = q.drawForLocked(t.Difficulty, q.Rounds)
		}
	}
}
