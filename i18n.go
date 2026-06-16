package main

import "strings"

// Supported quiz-content languages. Only question content (text, options,
// facts) is translated — the landing page and instructional UI stay English.
const (
	LangEN = "en"
	LangIS = "is"
	LangSV = "sv"
)

// langMeta describes a selectable language for the create form.
type langMeta struct {
	Code  string
	Label string
}

// languages is the ordered list shown in the language selector. English first
// because it is the default and the fallback for any untranslated content.
var languages = []langMeta{
	{LangEN, "English"},
	{LangIS, "Íslenska"},
	{LangSV, "Svenska"},
}

// normalizeLang coerces arbitrary input to a supported language, defaulting to
// English.
func normalizeLang(lang string) string {
	switch lang {
	case LangIS, LangSV:
		return lang
	default:
		return LangEN
	}
}

// QuestionL10n is a translated version of a question's content. Options must be
// in the same order as the English original so the Answer index still applies.
type QuestionL10n struct {
	Text    string
	Options []string
	Fact    string
}

// translationFor returns the translation of an English question Text for the
// given language, if one exists.
func translationFor(text, lang string) (QuestionL10n, bool) {
	switch lang {
	case LangIS:
		l, ok := translationsIS[text]
		return l, ok
	case LangSV:
		l, ok := translationsSV[text]
		return l, ok
	}
	return QuestionL10n{}, false
}

// localize resolves a question's content into the requested language. Any field
// (or the whole question) without a translation falls back to English — which
// is exactly the desired behaviour for disputed terms such as "hat trick".
func localize(q Question, lang string) (text string, options []string, fact string) {
	text, options, fact = q.Text, q.Options, q.Fact
	if lang == LangEN || lang == "" {
		return
	}
	l, ok := translationFor(q.Text, lang)
	if !ok {
		return
	}
	if strings.TrimSpace(l.Text) != "" {
		text = l.Text
	}
	if len(l.Options) == len(q.Options) {
		options = l.Options
	}
	if strings.TrimSpace(l.Fact) != "" {
		fact = l.Fact
	}
	return
}
