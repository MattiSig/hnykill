package main

import (
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

// captchaAnswer solves a generated prompt ("What is A + B?") the way a human
// would, so tests can submit the correct answer.
func captchaAnswer(t *testing.T, prompt string) string {
	t.Helper()
	var a, b int
	if _, err := fmt.Sscanf(prompt, "What is %d + %d?", &a, &b); err != nil {
		t.Fatalf("unexpected prompt %q: %v", prompt, err)
	}
	return strconv.Itoa(a + b)
}

func TestCaptchaRoundTrip(t *testing.T) {
	prompt, token := newCaptcha()
	answer := captchaAnswer(t, prompt)

	if !verifyCaptcha(token, answer) {
		t.Errorf("valid captcha rejected (prompt %q answer %q)", prompt, answer)
	}
	if verifyCaptcha(token, answer+"9") {
		t.Error("wrong answer accepted")
	}
	if verifyCaptcha(token+"x", answer) {
		t.Error("tampered token accepted")
	}
	if verifyCaptcha("not-a-token", answer) {
		t.Error("garbage token accepted")
	}
}

func TestCaptchaExpiry(t *testing.T) {
	// Forge an expired but correctly-signed token; it must still be rejected.
	payload := fmt.Sprintf("%d|7", time.Now().Add(-time.Minute).Unix())
	token := base64.RawURLEncoding.EncodeToString([]byte(payload)) + "." + sign(payload)
	if verifyCaptcha(token, "7") {
		t.Error("expired captcha accepted")
	}
}

func TestSanitizeSubmission(t *testing.T) {
	cases := []struct {
		name    string
		text    string
		opts    []string
		answer  int
		diff    string
		lang    string
		wantErr bool
		wantAns int // expected normalized answer index when valid
	}{
		{"ok", "Q?", []string{"A", "B", "C", ""}, 1, DiffAdult, "en", false, 1},
		{"blank text", "  ", []string{"A", "B"}, 0, DiffKids, "en", true, 0},
		{"too few options", "Q?", []string{"A", "", "", ""}, 0, DiffKids, "en", true, 0},
		{"answer on blank", "Q?", []string{"A", "", "C", ""}, 1, DiffKids, "en", true, 0},
		{"answer reindexes", "Q?", []string{"", "A", "B", ""}, 2, DiffNerd, "is", false, 1},
		{"bad difficulty", "Q?", []string{"A", "B"}, 0, "guru", "en", true, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			q, msg := sanitizeSubmission(c.text, c.opts, c.answer, "", c.diff, c.lang, "")
			if c.wantErr {
				if msg == "" {
					t.Fatalf("expected error, got valid question %+v", q)
				}
				return
			}
			if msg != "" {
				t.Fatalf("unexpected error: %s", msg)
			}
			if q.Answer != c.wantAns {
				t.Errorf("answer index = %d, want %d (options %v)", q.Answer, c.wantAns, q.Options)
			}
			if q.Options[q.Answer] == "" {
				t.Error("correct option is blank after sanitize")
			}
		})
	}
}

func TestFileStoreVoteAndOrder(t *testing.T) {
	path := filepath.Join(t.TempDir(), "community.json")
	s, err := newFileCommunityStore(path)
	if err != nil {
		t.Fatal(err)
	}

	q1 := &CommunityQuestion{Text: "first", Options: []string{"a", "b"}, Difficulty: DiffKids, Language: "en"}
	q2 := &CommunityQuestion{Text: "second", Options: []string{"a", "b"}, Difficulty: DiffKids, Language: "en"}
	if err := s.Add(q1); err != nil {
		t.Fatal(err)
	}
	if err := s.Add(q2); err != nil {
		t.Fatal(err)
	}

	// Vote q2 up twice from the same voter -> counts once.
	if voted, _ := s.Vote(q2.ID, "voter-A"); !voted {
		t.Error("first vote should count")
	}
	if voted, _ := s.Vote(q2.ID, "voter-A"); voted {
		t.Error("duplicate vote should not count")
	}
	if voted, _ := s.Vote(q2.ID, "voter-B"); !voted {
		t.Error("second distinct voter should count")
	}

	list, err := s.List("", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 2 {
		t.Fatalf("expected 2 questions, got %d", len(list))
	}
	if list[0].ID != q2.ID {
		t.Errorf("expected most-voted (%s) first, got %s", q2.ID, list[0].ID)
	}
	if list[0].Votes != 2 {
		t.Errorf("expected 2 votes, got %d", list[0].Votes)
	}

	// Reload from disk: votes and voters must persist.
	s2, err := newFileCommunityStore(path)
	if err != nil {
		t.Fatal(err)
	}
	reloaded, _ := s2.List("", "")
	if len(reloaded) != 2 || reloaded[0].Votes != 2 {
		t.Fatalf("persistence broken: %+v", reloaded)
	}
	if voted, _ := s2.Vote(q2.ID, "voter-A"); voted {
		t.Error("voter dedup not persisted across reload")
	}
}

func TestSeedIsIdempotentAndPreservesVotes(t *testing.T) {
	path := filepath.Join(t.TempDir(), "community.json")
	s, err := newFileCommunityStore(path)
	if err != nil {
		t.Fatal(err)
	}

	seed := officialQuestions()
	want := 84 * len(languages) // every built-in question in every language
	if len(seed) != want {
		t.Fatalf("officialQuestions() = %d, want %d", len(seed), want)
	}

	added, err := s.Seed(seed)
	if err != nil {
		t.Fatal(err)
	}
	if added != want || s.Count() != want {
		t.Fatalf("first seed added %d / count %d, want %d", added, s.Count(), want)
	}

	// Cast a vote on an official question.
	official, _ := s.List("en", "")
	if len(official) == 0 {
		t.Fatal("no english questions after seed")
	}
	if voted, _ := s.Vote(official[0].ID, "voter-A"); !voted {
		t.Fatal("vote on official question should count")
	}

	// Re-seed (simulates a restart/redeploy): nothing new, no duplicates.
	added2, err := s.Seed(officialQuestions())
	if err != nil {
		t.Fatal(err)
	}
	if added2 != 0 {
		t.Errorf("re-seed added %d, want 0", added2)
	}
	if s.Count() != want {
		t.Errorf("count after re-seed = %d, want %d", s.Count(), want)
	}

	// The vote must survive the re-seed.
	again, _ := s.List("en", "")
	var votes int
	for _, q := range again {
		if q.ID == official[0].ID {
			votes = q.Votes
		}
	}
	if votes != 1 {
		t.Errorf("vote lost across re-seed: votes = %d, want 1", votes)
	}

	// Stable, namespaced ids: same text, different language -> different id.
	if officialID("en", "x") == officialID("is", "x") {
		t.Error("officialID should differ by language")
	}
	if officialID("en", "x") != officialID("en", "x") {
		t.Error("officialID should be deterministic")
	}
}

func TestFileStoreFilter(t *testing.T) {
	s, _ := newFileCommunityStore(filepath.Join(t.TempDir(), "c.json"))
	_ = s.Add(&CommunityQuestion{Text: "en-kids", Options: []string{"a", "b"}, Difficulty: DiffKids, Language: "en"})
	_ = s.Add(&CommunityQuestion{Text: "is-nerd", Options: []string{"a", "b"}, Difficulty: DiffNerd, Language: "is"})

	if got, _ := s.List("is", ""); len(got) != 1 || got[0].Text != "is-nerd" {
		t.Errorf("language filter failed: %+v", got)
	}
	if got, _ := s.List("", DiffKids); len(got) != 1 || got[0].Text != "en-kids" {
		t.Errorf("difficulty filter failed: %+v", got)
	}
	if got, _ := s.List("en", DiffNerd); len(got) != 0 {
		t.Errorf("combined filter should be empty, got %+v", got)
	}
}
