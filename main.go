package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

//go:embed templates/*.html
var templateFS embed.FS

var templates = template.Must(template.ParseFS(templateFS, "templates/*.html"))

var store = NewStore()

var community = newCommunityStore()

func main() {
	seedOfficialQuestions(community)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleHome)
	mux.HandleFunc("POST /create", handleCreate)
	mux.HandleFunc("GET /admin/{id}", handleAdmin)
	mux.HandleFunc("GET /qr/{id}", handleQR)
	mux.HandleFunc("GET /j/{id}", handleJoin)
	mux.HandleFunc("POST /j/{id}", handleJoinSubmit)
	mux.HandleFunc("GET /play/{id}", handlePlay)
	mux.HandleFunc("GET /api/state/{id}", handleState)
	mux.HandleFunc("POST /api/answer/{id}", handleAnswer)
	mux.HandleFunc("POST /api/admin/{id}/{action}", handleAdminAction)
	mux.HandleFunc("GET /community", handleCommunityBrowse)
	mux.HandleFunc("GET /community/new", handleCommunityNew)
	mux.HandleFunc("POST /community/new", handleCommunitySubmit)
	mux.HandleFunc("POST /api/community/{id}/vote", handleCommunityVote)
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Printf("World Cup 2026 Quiz listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// ---- helpers ----

func baseURL(r *http.Request) string {
	scheme := "http"
	if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		scheme = proto
	} else if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}

func render(w http.ResponseWriter, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.ExecuteTemplate(w, name, data); err != nil {
		log.Printf("template %s: %v", name, err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func playerCookieName(quizID string) string { return "hny_pid_" + quizID }

// voterCookieName is the anonymous identity used to dedupe community upvotes
// (one vote per question per browser).
const voterCookieName = "hny_voter"

// ensureVoterID returns the requester's voter id, minting and setting a cookie
// if they don't have one yet.
func ensureVoterID(w http.ResponseWriter, r *http.Request) string {
	if c, err := r.Cookie(voterCookieName); err == nil && c.Value != "" {
		return c.Value
	}
	id := randString(hexAlphabet, 16)
	http.SetCookie(w, &http.Cookie{
		Name:     voterCookieName,
		Value:    id,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 365,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return id
}

// quizMeta returns display fields (title, emoji, accent) for a quiz, handling
// the synthetic "mixed" difficulty which has no entry in the difficulties map.
func quizMeta(q *Quiz) (title, emoji, accent string) {
	if q.Mixed {
		return "Mixed", "🎲", "violet"
	}
	m := difficulties[q.Difficulty]
	return m.Title, m.Emoji, m.Accent
}

// ---- handlers ----

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	type diffView struct {
		Key, Title, Tagline, Emoji, Accent string
		Points, Count                      int
	}
	var diffs []diffView
	for _, key := range difficultyOrder {
		d := difficulties[key]
		diffs = append(diffs, diffView{
			Key: d.Key, Title: d.Title, Tagline: d.Tagline, Emoji: d.Emoji,
			Accent: d.Accent, Points: d.Points, Count: len(d.Questions),
		})
	}
	render(w, "home.html", map[string]any{
		"Difficulties": diffs,
	})
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	diff := r.FormValue("difficulty")
	if diff != DiffMixed {
		if _, ok := difficulties[diff]; !ok {
			http.Error(w, "unknown difficulty", http.StatusBadRequest)
			return
		}
	}
	q := store.CreateQuiz(diff, r.FormValue("lang"))
	http.Redirect(w, r, "/admin/"+q.ID+"?t="+q.AdminToken, http.StatusSeeOther)
}

func handleAdmin(w http.ResponseWriter, r *http.Request) {
	q, ok := store.Get(r.PathValue("id"))
	if !ok {
		notFoundPage(w, r)
		return
	}
	if r.URL.Query().Get("t") != q.AdminToken {
		render(w, "denied.html", nil)
		return
	}
	title, emoji, accent := quizMeta(q)
	total := len(q.Questions)
	if q.Mixed {
		total = q.Rounds
	}
	// Difficulty choices offered for per-team assignment in mixed mode.
	type diffOpt struct{ Key, Title, Emoji string }
	var diffOpts []diffOpt
	for _, key := range difficultyOrder {
		d := difficulties[key]
		diffOpts = append(diffOpts, diffOpt{d.Key, d.Title, d.Emoji})
	}
	render(w, "admin.html", map[string]any{
		"QuizID":      q.ID,
		"Token":       q.AdminToken,
		"Difficulty":  title,
		"Emoji":       emoji,
		"Accent":      accent,
		"Total":       total,
		"Points":      q.Points,
		"Mixed":       q.Mixed,
		"DiffOptions": diffOpts,
		"Lang":        q.Lang,
		"Languages":   languages,
		"JoinURL":     baseURL(r) + "/j/" + q.ID,
	})
}

func handleQR(w http.ResponseWriter, r *http.Request) {
	q, ok := store.Get(r.PathValue("id"))
	if !ok {
		http.NotFound(w, r)
		return
	}
	joinURL := baseURL(r) + "/j/" + q.ID
	png, err := qrcode.Encode(joinURL, qrcode.Medium, 320)
	if err != nil {
		http.Error(w, "qr error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "no-store")
	w.Write(png)
}

func handleJoin(w http.ResponseWriter, r *http.Request) {
	q, ok := store.Get(r.PathValue("id"))
	if !ok {
		notFoundPage(w, r)
		return
	}
	// Already joined? Jump straight to play.
	if c, err := r.Cookie(playerCookieName(q.ID)); err == nil {
		if _, ok := q.Team(c.Value); ok {
			http.Redirect(w, r, "/play/"+q.ID, http.StatusSeeOther)
			return
		}
	}
	title, emoji, accent := quizMeta(q)
	render(w, "join.html", map[string]any{
		"QuizID":     q.ID,
		"Difficulty": title,
		"Emoji":      emoji,
		"Accent":     accent,
	})
}

func handleJoinSubmit(w http.ResponseWriter, r *http.Request) {
	q, ok := store.Get(r.PathValue("id"))
	if !ok {
		notFoundPage(w, r)
		return
	}
	name := strings.TrimSpace(r.FormValue("name"))
	if name == "" {
		name = "Team " + randString(idAlphabet, 3)
	}
	if len(name) > 30 {
		name = name[:30]
	}
	t := q.AddTeam(name)
	http.SetCookie(w, &http.Cookie{
		Name:     playerCookieName(q.ID),
		Value:    t.ID,
		Path:     "/",
		MaxAge:   60 * 60 * 24,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, "/play/"+q.ID, http.StatusSeeOther)
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	q, ok := store.Get(r.PathValue("id"))
	if !ok {
		notFoundPage(w, r)
		return
	}
	c, err := r.Cookie(playerCookieName(q.ID))
	if err != nil {
		http.Redirect(w, r, "/j/"+q.ID, http.StatusSeeOther)
		return
	}
	t, ok := q.Team(c.Value)
	if !ok {
		http.Redirect(w, r, "/j/"+q.ID, http.StatusSeeOther)
		return
	}
	title, emoji, accent := quizMeta(q)
	// In mixed mode show the team's own assigned difficulty.
	if q.Mixed {
		if m, ok := difficulties[t.Difficulty]; ok {
			title, emoji, accent = m.Title, m.Emoji, m.Accent
		}
	}
	render(w, "play.html", map[string]any{
		"QuizID":     q.ID,
		"TeamName":   t.Name,
		"Difficulty": title,
		"Emoji":      emoji,
		"Accent":     accent,
	})
}

func handleState(w http.ResponseWriter, r *http.Request) {
	q, ok := store.Get(r.PathValue("id"))
	if !ok {
		http.NotFound(w, r)
		return
	}

	isAdmin := r.URL.Query().Get("t") == q.AdminToken
	var teamID string
	if c, err := r.Cookie(playerCookieName(q.ID)); err == nil {
		teamID = c.Value
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	revealed := q.Phase == PhaseRevealed
	finished := q.Phase == PhaseFinished
	title, _, _ := quizMeta(q)

	var viewer *Team
	if t, ok := q.teams[teamID]; ok {
		viewer = t
	}

	state := map[string]any{
		"phase":         q.Phase,
		"current":       q.Current,
		"total":         q.totalLocked(),
		"points":        q.Points,
		"isAdmin":       isAdmin,
		"mixed":         q.Mixed,
		"lang":          q.Lang,
		"difficulty":    title,
		"teamCount":     len(q.teams),
		"answeredCount": 0,
		"teams":         q.leaderboardLocked(true, revealed || finished),
	}

	// Current question payload (visible during question/revealed phases). In
	// mixed mode the question is per-team, so the admin (who has no single
	// question) sees only progress while each player sees their own question.
	if q.Phase == PhaseQuestion || q.Phase == PhaseRevealed {
		if que, ok := q.questionForLocked(viewer); ok && !(q.Mixed && viewer == nil) {
			text, options, factText := localize(que, q.Lang)
			answer := -1
			fact := ""
			if revealed || isAdmin {
				answer = que.Answer
			}
			if revealed {
				fact = factText
			}
			state["question"] = text
			state["options"] = options
			state["answer"] = answer
			state["fact"] = fact
		}

		answered := 0
		for _, t := range q.teams {
			if _, ok := t.answers[q.Current]; ok {
				answered++
			}
		}
		state["answeredCount"] = answered
	}

	// The requesting player's own answer for the current question.
	yourChoice := -1
	yourScore := 0
	if t, ok := q.teams[teamID]; ok {
		yourScore = t.Score
		if q.Phase == PhaseQuestion || q.Phase == PhaseRevealed {
			if c, ok := t.answers[q.Current]; ok {
				yourChoice = c
			}
		}
		state["youJoined"] = true
		state["yourName"] = t.Name
	}
	state["yourChoice"] = yourChoice
	state["yourScore"] = yourScore

	writeJSON(w, state)
}

func handleAnswer(w http.ResponseWriter, r *http.Request) {
	q, ok := store.Get(r.PathValue("id"))
	if !ok {
		http.NotFound(w, r)
		return
	}
	c, err := r.Cookie(playerCookieName(q.ID))
	if err != nil {
		http.Error(w, "not joined", http.StatusForbidden)
		return
	}
	choice, _ := strconv.Atoi(r.FormValue("choice"))
	ok = q.SubmitAnswer(c.Value, choice)
	writeJSON(w, map[string]any{"ok": ok})
}

func handleAdminAction(w http.ResponseWriter, r *http.Request) {
	q, ok := store.Get(r.PathValue("id"))
	if !ok {
		http.NotFound(w, r)
		return
	}
	if r.URL.Query().Get("t") != q.AdminToken {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	switch r.PathValue("action") {
	case "start":
		q.Start()
	case "reveal":
		q.Reveal()
	case "next":
		q.Next()
	case "reset":
		q.Reset()
	case "adjust":
		delta, _ := strconv.Atoi(r.FormValue("delta"))
		q.Adjust(r.FormValue("team"), delta)
	case "setdiff":
		q.SetTeamDifficulty(r.FormValue("team"), r.FormValue("difficulty"))
	case "setlang":
		q.SetLang(r.FormValue("lang"))
	default:
		http.Error(w, "unknown action", http.StatusBadRequest)
		return
	}
	writeJSON(w, map[string]any{"ok": true})
}

// ---- community questions ----

// diffOption / langOption are small view models for the submit form selects and
// the browse-page filters.
type diffOption struct{ Key, Title, Emoji string }
type langOption struct{ Code, Label string }

func difficultyOptions() []diffOption {
	var out []diffOption
	for _, key := range difficultyOrder {
		d := difficulties[key]
		out = append(out, diffOption{d.Key, d.Title, d.Emoji})
	}
	return out
}

func languageOptions() []langOption {
	var out []langOption
	for _, l := range languages {
		out = append(out, langOption{l.Code, l.Label})
	}
	return out
}

// communityCardView is a question prepared for display: the correct option is
// flagged and the difficulty/language get human labels.
type communityCardView struct {
	ID         string
	Text       string
	Options    []optionView
	Fact       string
	Difficulty string
	Language   string
	Author     string
	Official   bool
	Votes      int
	Voted      bool
}

type optionView struct {
	Text    string
	Correct bool
}

func langLabel(code string) string {
	for _, l := range languages {
		if l.Code == code {
			return l.Label
		}
	}
	return code
}

func diffLabel(key string) string {
	if d, ok := difficulties[key]; ok {
		return d.Emoji + " " + d.Title
	}
	return key
}

func handleCommunityBrowse(w http.ResponseWriter, r *http.Request) {
	voterID := ensureVoterID(w, r)

	lang := r.URL.Query().Get("lang")
	if lang != "" {
		lang = normalizeLang(lang)
	}
	difficulty := r.URL.Query().Get("difficulty")
	switch difficulty {
	case DiffKids, DiffAdult, DiffNerd:
	default:
		difficulty = ""
	}

	items, err := community.List(lang, difficulty)
	if err != nil {
		log.Printf("community list: %v", err)
		http.Error(w, "could not load questions", http.StatusInternalServerError)
		return
	}

	cards := make([]communityCardView, 0, len(items))
	for _, q := range items {
		opts := make([]optionView, len(q.Options))
		for i, o := range q.Options {
			opts[i] = optionView{Text: o, Correct: i == q.Answer}
		}
		cards = append(cards, communityCardView{
			ID:         q.ID,
			Text:       q.Text,
			Options:    opts,
			Fact:       q.Fact,
			Difficulty: diffLabel(q.Difficulty),
			Language:   langLabel(q.Language),
			Author:     q.Author,
			Official:   q.Author == officialAuthor,
			Votes:      q.Votes,
			Voted:      q.Voters[voterID],
		})
	}

	render(w, "community.html", map[string]any{
		"Cards":        cards,
		"Count":        len(cards),
		"Total":        community.Count(),
		"Languages":    languageOptions(),
		"Difficulties": difficultyOptions(),
		"FilterLang":   lang,
		"FilterDiff":   difficulty,
		"Submitted":    r.URL.Query().Get("submitted") == "1",
	})
}

func handleCommunityNew(w http.ResponseWriter, r *http.Request) {
	renderCommunityForm(w, "", nil)
}

// renderCommunityForm shows the submission form with a fresh captcha. errMsg and
// prev (the previously entered values) let it re-render after a failed submit
// without losing the user's input.
func renderCommunityForm(w http.ResponseWriter, errMsg string, prev map[string]any) {
	// Normalize so the template never has to guard nil/short values.
	if prev == nil {
		prev = map[string]any{}
	}
	opts, _ := prev["Options"].([]string)
	for len(opts) < maxOptions {
		opts = append(opts, "")
	}
	prev["Options"] = opts[:maxOptions]
	if _, ok := prev["Answer"].(int); !ok {
		prev["Answer"] = -1
	}
	for _, k := range []string{"Text", "Fact", "Difficulty", "Language", "Author"} {
		if _, ok := prev[k].(string); !ok {
			prev[k] = ""
		}
	}

	prompt, token := newCaptcha()
	data := map[string]any{
		"Error":         errMsg,
		"CaptchaPrompt": prompt,
		"CaptchaToken":  token,
		"Languages":     languageOptions(),
		"Difficulties":  difficultyOptions(),
		"Prev":          prev,
	}
	render(w, "community_new.html", data)
}

func handleCommunitySubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		renderCommunityForm(w, "Could not read the form. Please try again.", nil)
		return
	}

	answerIdx, _ := strconv.Atoi(r.FormValue("answer"))
	rawOptions := []string{
		r.FormValue("option0"), r.FormValue("option1"),
		r.FormValue("option2"), r.FormValue("option3"),
	}
	prev := map[string]any{
		"Text":       r.FormValue("text"),
		"Options":    rawOptions,
		"Answer":     answerIdx,
		"Fact":       r.FormValue("fact"),
		"Difficulty": r.FormValue("difficulty"),
		"Language":   r.FormValue("language"),
		"Author":     r.FormValue("author"),
	}

	if !verifyCaptcha(r.FormValue("captcha_token"), r.FormValue("captcha_answer")) {
		renderCommunityForm(w, "Captcha was wrong or expired — please solve the new one.", prev)
		return
	}

	q, msg := sanitizeSubmission(
		r.FormValue("text"), rawOptions, answerIdx,
		r.FormValue("fact"), r.FormValue("difficulty"), r.FormValue("language"), r.FormValue("author"),
	)
	if msg != "" {
		renderCommunityForm(w, msg, prev)
		return
	}

	if err := community.Add(q); err != nil {
		log.Printf("community add: %v", err)
		renderCommunityForm(w, "Something went wrong saving your question. Please try again.", prev)
		return
	}
	http.Redirect(w, r, "/community?submitted=1", http.StatusSeeOther)
}

func handleCommunityVote(w http.ResponseWriter, r *http.Request) {
	voterID := ensureVoterID(w, r)
	id := r.PathValue("id")
	voted, err := community.Vote(id, voterID)
	if err != nil {
		log.Printf("community vote: %v", err)
		http.Error(w, "vote failed", http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]any{"ok": true, "counted": voted})
}

func notFoundPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	render(w, "notfound.html", nil)
}
