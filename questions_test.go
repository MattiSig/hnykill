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

// TestTranslations ensures every question is fully translated into each
// non-English language, with the option count preserved so the Answer index
// stays valid. Missing or malformed translations would silently fall back to
// English at runtime, so we guard against that here.
func TestTranslations(t *testing.T) {
	for _, lang := range []string{LangIS, LangSV} {
		for _, key := range difficultyOrder {
			for i, q := range difficulties[key].Questions {
				where := lang + " " + key + " Q" + itoa(i)
				l, ok := translationFor(q.Text, lang)
				if !ok {
					t.Errorf("%s: missing translation for %q", where, q.Text)
					continue
				}
				if strings.TrimSpace(l.Text) == "" {
					t.Errorf("%s: empty translated text", where)
				}
				if strings.TrimSpace(l.Fact) == "" {
					t.Errorf("%s: empty translated fact", where)
				}
				if len(l.Options) != len(q.Options) {
					t.Errorf("%s: translated options count %d != original %d", where, len(l.Options), len(q.Options))
					continue
				}
				for j, o := range l.Options {
					if strings.TrimSpace(o) == "" {
						t.Errorf("%s: empty translated option %d", where, j)
					}
				}
			}
		}
	}
}

// TestLocalizeFallback verifies that an untranslated question falls back to
// English content (the desired behaviour for disputed terms).
func TestLocalizeFallback(t *testing.T) {
	q := Question{Text: "____no such question____", Options: []string{"a", "b"}, Answer: 0, Fact: "f"}
	text, opts, fact := localize(q, LangIS)
	if text != q.Text || fact != q.Fact || len(opts) != len(q.Options) {
		t.Errorf("expected English fallback, got text=%q fact=%q opts=%v", text, fact, opts)
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
