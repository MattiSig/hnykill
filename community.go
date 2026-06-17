package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	mrand "math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// CommunityQuestion is a question submitted by a real user for the community
// board. Unlike the built-in bank (which is translated into every language), a
// community question is authored in a SINGLE language and shown as-is.
type CommunityQuestion struct {
	ID         string          `json:"id"`
	Text       string          `json:"text"`
	Options    []string        `json:"options"`
	Answer     int             `json:"answer"` // index into Options of the correct option
	Fact       string          `json:"fact,omitempty"`
	Difficulty string          `json:"difficulty"` // kids | adult | nerd
	Language   string          `json:"language"`   // en | is | sv
	Author     string          `json:"author,omitempty"`
	Votes      int             `json:"votes"`
	CreatedAt  time.Time       `json:"created_at"`
	Voters     map[string]bool `json:"voters,omitempty"` // voter ids that have upvoted (file backend)
}

// CommunityStore persists community questions and their vote tallies. There are
// two implementations: a JSON file (local/dev) and Postgres (production).
type CommunityStore interface {
	// Add stores a new question. It assigns ID, Votes and CreatedAt if unset.
	Add(q *CommunityQuestion) error
	// List returns questions sorted by votes (desc) then newest first, optionally
	// filtered by language and/or difficulty (empty string means "any").
	List(lang, difficulty string) ([]*CommunityQuestion, error)
	// Vote records an upvote from voterID. It returns true if the vote counted
	// (false if that voter had already voted for this question).
	Vote(id, voterID string) (bool, error)
	// Count returns the total number of stored questions.
	Count() int
	// Seed inserts the given questions if they are not already present (matched
	// by ID). It is idempotent so it can run on every startup without creating
	// duplicates or disturbing existing vote tallies. Returns how many were added.
	Seed(qs []*CommunityQuestion) (int, error)
}

// officialAuthor labels the starter questions seeded from the built-in bank so
// the UI can distinguish them from genuine community submissions.
const officialAuthor = "Official"

// officialID derives a stable id from a seeded question's language and text, so
// re-seeding is idempotent (the same question always maps to the same row).
func officialID(lang, text string) string {
	sum := sha256.Sum256([]byte(lang + "\x00" + text))
	return hex.EncodeToString(sum[:])[:12]
}

// officialQuestions expands the built-in question bank into one single-language
// community question per (question, language), translated via localize so each
// card reads naturally in its language.
func officialQuestions() []*CommunityQuestion {
	var out []*CommunityQuestion
	for _, key := range difficultyOrder {
		meta := difficulties[key]
		for _, q := range meta.Questions {
			for _, lang := range languages {
				text, options, fact := localize(q, lang.Code)
				out = append(out, &CommunityQuestion{
					ID:         officialID(lang.Code, text),
					Text:       text,
					Options:    options,
					Answer:     q.Answer,
					Fact:       fact,
					Difficulty: key,
					Language:   lang.Code,
					Author:     officialAuthor,
				})
			}
		}
	}
	return out
}

// drawQuestions returns up to n play questions for the chosen source. The
// official source deals a fresh random subset of the built-in bank (replayable).
// The community source plays the crowd's favourites: the top-voted community
// questions for this language and difficulty, lightly shuffled for variety, so
// votes actually shape the game. If the community pool is empty it falls back to
// the built-in bank so a game is never left without questions.
func drawQuestions(source, lang, difficulty string, n int) []Question {
	if source == SourceCommunity {
		if items, err := community.List(lang, difficulty); err == nil && len(items) > 0 {
			// items are sorted by votes desc; keep the better-voted top slice,
			// then shuffle within it so repeat games aren't identical.
			if n > 0 && len(items) > n*2 {
				items = items[:n*2]
			}
			qs := make([]Question, 0, len(items))
			for _, c := range items {
				qs = append(qs, Question{Text: c.Text, Options: c.Options, Answer: c.Answer, Fact: c.Fact})
			}
			mrand.Shuffle(len(qs), func(i, j int) { qs[i], qs[j] = qs[j], qs[i] })
			if n > 0 && n < len(qs) {
				qs = qs[:n]
			}
			return qs
		}
	}
	return pickN(difficulties[difficulty], n)
}

// seedOfficialQuestions loads the built-in bank into the community board. Safe to
// call on every startup.
func seedOfficialQuestions(store CommunityStore) {
	added, err := store.Seed(officialQuestions())
	if err != nil {
		log.Printf("community seed: %v", err)
		return
	}
	if added > 0 {
		log.Printf("community seed: added %d official starter questions", added)
	}
}

// newCommunityStore picks Postgres when DATABASE_URL is set (production on
// Railway) and otherwise a JSON file under DATA_DIR (default ./data) for local
// development. It never returns nil: on failure it logs and falls back to file.
func newCommunityStore() CommunityStore {
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		s, err := newPGCommunityStore(dsn)
		if err == nil {
			log.Printf("community store: postgres")
			return s
		}
		log.Printf("community store: postgres unavailable (%v), falling back to file", err)
	}
	dir := os.Getenv("DATA_DIR")
	if dir == "" {
		dir = "data"
	}
	s, err := newFileCommunityStore(filepath.Join(dir, "community.json"))
	if err != nil {
		log.Printf("community store: file init failed (%v), using in-memory only", err)
	}
	log.Printf("community store: file (%s)", filepath.Join(dir, "community.json"))
	return s
}

// ---- validation ----

const (
	maxQuestionLen = 200
	maxOptionLen   = 80
	maxFactLen     = 240
	maxAuthorLen   = 30
	maxOptions     = 4
	minOptions     = 2
)

// sanitizeSubmission validates and normalizes raw form input into a
// CommunityQuestion. It returns a user-facing error message (not an error value)
// describing the first problem, or "" when the submission is valid.
func sanitizeSubmission(text string, rawOptions []string, answerIdx int, fact, difficulty, language, author string) (*CommunityQuestion, string) {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil, "Please enter the question text."
	}
	if len(text) > maxQuestionLen {
		return nil, fmt.Sprintf("Question is too long (max %d characters).", maxQuestionLen)
	}

	// Keep non-empty options, remembering their original index so the chosen
	// correct answer still points at the right one after blanks are dropped.
	var options []string
	correct := -1
	for i, o := range rawOptions {
		o = strings.TrimSpace(o)
		if o == "" {
			continue
		}
		if len(o) > maxOptionLen {
			return nil, fmt.Sprintf("Answer options must be %d characters or fewer.", maxOptionLen)
		}
		if i == answerIdx {
			correct = len(options)
		}
		options = append(options, o)
	}
	if len(options) < minOptions {
		return nil, fmt.Sprintf("Please provide at least %d answer options.", minOptions)
	}
	if len(options) > maxOptions {
		return nil, fmt.Sprintf("Please provide at most %d answer options.", maxOptions)
	}
	if correct < 0 {
		return nil, "Please mark which option is the correct answer."
	}

	fact = strings.TrimSpace(fact)
	if len(fact) > maxFactLen {
		return nil, fmt.Sprintf("Fun fact is too long (max %d characters).", maxFactLen)
	}

	switch difficulty {
	case DiffKids, DiffAdult, DiffNerd:
	default:
		return nil, "Please choose a difficulty."
	}

	author = strings.TrimSpace(author)
	if len(author) > maxAuthorLen {
		author = author[:maxAuthorLen]
	}

	return &CommunityQuestion{
		Text:       text,
		Options:    options,
		Answer:     correct,
		Fact:       fact,
		Difficulty: difficulty,
		Language:   normalizeLang(language),
		Author:     author,
	}, ""
}

// sortQuestions orders questions by votes (desc) then newest first. Used by both
// backends so ordering is identical regardless of storage.
func sortQuestions(qs []*CommunityQuestion) {
	sort.SliceStable(qs, func(i, j int) bool {
		if qs[i].Votes != qs[j].Votes {
			return qs[i].Votes > qs[j].Votes
		}
		return qs[i].CreatedAt.After(qs[j].CreatedAt)
	})
}

// ---- file backend ----

// fileCommunityStore keeps questions in memory and flushes the whole set to a
// JSON file after each mutation (atomic temp-file + rename). It matches the
// in-memory-map style of store.go and needs no external service.
type fileCommunityStore struct {
	mu    sync.Mutex
	path  string
	items []*CommunityQuestion
}

func newFileCommunityStore(path string) (*fileCommunityStore, error) {
	s := &fileCommunityStore{path: path}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return s, nil // fresh store
		}
		return s, err
	}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &s.items); err != nil {
			return s, fmt.Errorf("parse %s: %w", path, err)
		}
	}
	return s, nil
}

func (s *fileCommunityStore) Add(q *CommunityQuestion) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if q.ID == "" {
		q.ID = randString(hexAlphabet, 12)
	}
	if q.CreatedAt.IsZero() {
		q.CreatedAt = time.Now()
	}
	if q.Voters == nil {
		q.Voters = map[string]bool{}
	}
	s.items = append(s.items, q)
	return s.saveLocked()
}

func (s *fileCommunityStore) List(lang, difficulty string) ([]*CommunityQuestion, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]*CommunityQuestion, 0, len(s.items))
	for _, q := range s.items {
		if lang != "" && q.Language != lang {
			continue
		}
		if difficulty != "" && q.Difficulty != difficulty {
			continue
		}
		out = append(out, q)
	}
	sortQuestions(out)
	return out, nil
}

func (s *fileCommunityStore) Vote(id, voterID string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, q := range s.items {
		if q.ID != id {
			continue
		}
		if q.Voters == nil {
			q.Voters = map[string]bool{}
		}
		if q.Voters[voterID] {
			return false, nil // already voted
		}
		q.Voters[voterID] = true
		q.Votes++
		return true, s.saveLocked()
	}
	return false, nil // unknown question: treat as no-op
}

func (s *fileCommunityStore) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items)
}

func (s *fileCommunityStore) Seed(qs []*CommunityQuestion) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	existing := make(map[string]bool, len(s.items))
	for _, q := range s.items {
		existing[q.ID] = true
	}
	added := 0
	for _, q := range qs {
		if q.ID == "" || existing[q.ID] {
			continue
		}
		if q.CreatedAt.IsZero() {
			q.CreatedAt = time.Now()
		}
		if q.Voters == nil {
			q.Voters = map[string]bool{}
		}
		s.items = append(s.items, q)
		existing[q.ID] = true
		added++
	}
	if added == 0 {
		return 0, nil
	}
	return added, s.saveLocked()
}

// saveLocked atomically persists the current items. Caller must hold s.mu.
func (s *fileCommunityStore) saveLocked() error {
	if s.path == "" {
		return nil // in-memory only
	}
	dir := filepath.Dir(s.path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s.items, "", "  ")
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(dir, "community-*.tmp")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		os.Remove(tmpName)
		return err
	}
	if err := tmp.Close(); err != nil {
		os.Remove(tmpName)
		return err
	}
	return os.Rename(tmpName, s.path)
}

// ---- postgres backend ----

// pgCommunityStore persists to Railway Postgres. Pure-Go pgx keeps the build
// CGO-free (the Dockerfile sets CGO_ENABLED=0).
type pgCommunityStore struct {
	pool *pgxpool.Pool
}

func newPGCommunityStore(dsn string) (*pgCommunityStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	schema := `
CREATE TABLE IF NOT EXISTS community_questions (
    id         text PRIMARY KEY,
    text       text NOT NULL,
    options    jsonb NOT NULL,
    answer     int NOT NULL,
    fact       text NOT NULL DEFAULT '',
    difficulty text NOT NULL,
    language   text NOT NULL,
    author     text NOT NULL DEFAULT '',
    votes      int NOT NULL DEFAULT 0,
    created_at timestamptz NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS community_votes (
    question_id text NOT NULL,
    voter_id    text NOT NULL,
    PRIMARY KEY (question_id, voter_id)
);`
	if _, err := pool.Exec(ctx, schema); err != nil {
		pool.Close()
		return nil, err
	}
	return &pgCommunityStore{pool: pool}, nil
}

func (s *pgCommunityStore) Add(q *CommunityQuestion) error {
	if q.ID == "" {
		q.ID = randString(hexAlphabet, 12)
	}
	if q.CreatedAt.IsZero() {
		q.CreatedAt = time.Now()
	}
	opts, err := json.Marshal(q.Options)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = s.pool.Exec(ctx, `
INSERT INTO community_questions (id, text, options, answer, fact, difficulty, language, author, votes, created_at)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
		q.ID, q.Text, opts, q.Answer, q.Fact, q.Difficulty, q.Language, q.Author, q.Votes, q.CreatedAt)
	return err
}

func (s *pgCommunityStore) List(lang, difficulty string) ([]*CommunityQuestion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := s.pool.Query(ctx, `
SELECT id, text, options, answer, fact, difficulty, language, author, votes, created_at
FROM community_questions
WHERE ($1 = '' OR language = $1) AND ($2 = '' OR difficulty = $2)
ORDER BY votes DESC, created_at DESC`, lang, difficulty)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []*CommunityQuestion
	for rows.Next() {
		var (
			q    CommunityQuestion
			opts []byte
		)
		if err := rows.Scan(&q.ID, &q.Text, &opts, &q.Answer, &q.Fact, &q.Difficulty, &q.Language, &q.Author, &q.Votes, &q.CreatedAt); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(opts, &q.Options); err != nil {
			return nil, err
		}
		qq := q
		out = append(out, &qq)
	}
	return out, rows.Err()
}

func (s *pgCommunityStore) Vote(id, voterID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return false, err
	}
	defer tx.Rollback(ctx)

	tag, err := tx.Exec(ctx, `INSERT INTO community_votes (question_id, voter_id) VALUES ($1,$2) ON CONFLICT DO NOTHING`, id, voterID)
	if err != nil {
		return false, err
	}
	if tag.RowsAffected() == 0 {
		return false, nil // already voted
	}
	if _, err := tx.Exec(ctx, `UPDATE community_questions SET votes = votes + 1 WHERE id = $1`, id); err != nil {
		return false, err
	}
	return true, tx.Commit(ctx)
}

func (s *pgCommunityStore) Count() int {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var n int
	if err := s.pool.QueryRow(ctx, `SELECT count(*) FROM community_questions`).Scan(&n); err != nil {
		return 0
	}
	return n
}

func (s *pgCommunityStore) Seed(qs []*CommunityQuestion) (int, error) {
	if len(qs) == 0 {
		return 0, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	batch := &pgx.Batch{}
	for _, q := range qs {
		opts, err := json.Marshal(q.Options)
		if err != nil {
			return 0, err
		}
		created := q.CreatedAt
		if created.IsZero() {
			created = time.Now()
		}
		batch.Queue(`
INSERT INTO community_questions (id, text, options, answer, fact, difficulty, language, author, votes, created_at)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) ON CONFLICT (id) DO NOTHING`,
			q.ID, q.Text, opts, q.Answer, q.Fact, q.Difficulty, q.Language, q.Author, q.Votes, created)
	}
	br := s.pool.SendBatch(ctx, batch)
	defer br.Close()
	added := 0
	for range qs {
		tag, err := br.Exec()
		if err != nil {
			return added, err
		}
		added += int(tag.RowsAffected())
	}
	return added, nil
}
