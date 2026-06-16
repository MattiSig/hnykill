package main

import (
	"strings"
	"testing"
)

// TestQuestionBanks validates every curated question so a typo can never ship
// a broken or unanswerable quiz.
func TestQuestionBanks(t *testing.T) {
	for _, key := range difficultyOrder {
		meta, ok := difficulties[key]
		if !ok {
			t.Fatalf("difficulty %q missing from difficulties map", key)
		}
		if meta.PlayCount > len(meta.Questions) {
			t.Errorf("%s: PlayCount %d exceeds pool size %d", key, meta.PlayCount, len(meta.Questions))
		}
		if meta.PlayCount < 1 {
			t.Errorf("%s: PlayCount must be >= 1", key)
		}
		if meta.Points <= 0 {
			t.Errorf("%s: Points must be positive", key)
		}

		for i, q := range meta.Questions {
			where := key + " Q" + itoa(i)
			if strings.TrimSpace(q.Text) == "" {
				t.Errorf("%s: empty question text", where)
			}
			if len(q.Options) < 2 || len(q.Options) > 5 {
				t.Errorf("%s: expected 2-5 options, got %d", where, len(q.Options))
			}
			if q.Answer < 0 || q.Answer >= len(q.Options) {
				t.Errorf("%s: answer index %d out of range (0-%d)", where, q.Answer, len(q.Options)-1)
			}
			if strings.TrimSpace(q.Fact) == "" {
				t.Errorf("%s: missing reveal fact", where)
			}
			seen := map[string]bool{}
			for _, o := range q.Options {
				o = strings.TrimSpace(o)
				if o == "" {
					t.Errorf("%s: empty option", where)
				}
				if seen[strings.ToLower(o)] {
					t.Errorf("%s: duplicate option %q", where, o)
				}
				seen[strings.ToLower(o)] = true
			}
		}
	}
}

// TestPickQuestions ensures a game draws the right count without duplicates.
func TestPickQuestions(t *testing.T) {
	for _, key := range difficultyOrder {
		meta := difficulties[key]
		got := pickQuestions(meta)
		if len(got) != meta.PlayCount {
			t.Errorf("%s: pickQuestions returned %d, want %d", key, len(got), meta.PlayCount)
		}
		seen := map[string]bool{}
		for _, q := range got {
			if seen[q.Text] {
				t.Errorf("%s: pickQuestions returned duplicate question %q", key, q.Text)
			}
			seen[q.Text] = true
		}
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}
