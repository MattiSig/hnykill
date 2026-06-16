package main

import (
	"crypto/rand"
	"sort"
	"sync"
	"time"
)

// Phase represents the lifecycle stage of a quiz.
const (
	PhaseLobby     = "lobby"     // waiting for teams to join
	PhaseQuestion  = "question"  // a question is live, teams can answer
	PhaseRevealed  = "revealed"  // the answer is shown, scores updated
	PhaseFinished  = "finished"  // all questions done, final leaderboard
)

// Team is a group of attendees playing together, run by a captain.
type Team struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Score    int    `json:"score"`
	JoinedAt time.Time `json:"-"`
	// answers maps question index -> chosen option index for this team.
	answers map[int]int
}

// Quiz is a single live quiz session held entirely in memory.
type Quiz struct {
	ID         string
	AdminToken string
	Difficulty string
	Questions  []Question
	Points     int

	Phase      string
	Current    int // index of the current question (valid once started)

	teams map[string]*Team
	order []string // team ids in join order

	CreatedAt time.Time
	mu        sync.Mutex
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
	idAlphabet    = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789" // no ambiguous chars (0/O, 1/I)
	hexAlphabet   = "0123456789abcdef"
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

// CreateQuiz builds a new quiz for the given difficulty and returns it.
func (s *Store) CreateQuiz(difficulty string) *Quiz {
	meta, ok := difficulties[difficulty]
	if !ok {
		meta = difficulties[DiffAdult]
	}

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
		Difficulty: meta.Key,
		Questions:  meta.Questions,
		Points:     meta.Points,
		Phase:      PhaseLobby,
		Current:    0,
		teams:      make(map[string]*Team),
		CreatedAt:  time.Now(),
	}
	s.quizzes[id] = q
	return q
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
	q.teams[t.ID] = t
	q.order = append(q.order, t.ID)
	return t
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
	if choice < 0 || choice >= len(q.Questions[q.Current].Options) {
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
	ID       string `json:"id"`
	Name     string `json:"name"`
	Score    int    `json:"score"`
	Answered bool   `json:"answered"` // answered the current question
	Choice   int    `json:"choice"`   // -1 if not answered / not revealed
	Correct  bool   `json:"correct"`  // whether their answer was correct (revealed only)
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
				correct = c == q.Questions[q.Current].Answer
			}
		}
		views = append(views, teamView{
			ID:       t.ID,
			Name:     t.Name,
			Score:    t.Score,
			Answered: answered,
			Choice:   choice,
			Correct:  correct,
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
		q.Phase = PhaseQuestion
		q.Current = 0
	}
}

// Reveal scores the current question and shows the answer.
func (q *Quiz) Reveal() {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.Phase != PhaseQuestion {
		return
	}
	answer := q.Questions[q.Current].Answer
	for _, t := range q.teams {
		if c, ok := t.answers[q.Current]; ok && c == answer {
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
	if q.Current+1 >= len(q.Questions) {
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
	}
}
