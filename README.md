# ⚽ World Cup 2026 Quiz

A tiny, self-contained Go web app for running a **FIFA World Cup 2026** trivia
quiz in family settings. The quiz master mints a quiz, everyone else scans a QR
code to join as a team, and scores are kept live — Kahoot-style.

> Built for the first 48-team World Cup, co-hosted by **Canada, Mexico & the USA**.

## Features

- **Three difficulties** — `Kids` 🧒, `Adult` 🧑 and `Nerd` 🤓, each with a
  curated, fact-checked question bank about the 2026 tournament.
- **Mint & host** — whoever creates a quiz becomes the **admin / scorekeeper**.
- **QR code join** — the admin dashboard shows a QR code (and a 4-letter join
  code) for attendees to scan and join as team captains from their phones.
- **Live play** — teams lock in multiple-choice answers; the admin reveals the
  answer, scores update automatically, with a live scoreboard and final podium.
- **Manual score control** — the admin can nudge any team's score up/down.
- **Community questions** 🗳️ — anyone can submit their own questions (any
  difficulty, any single language) and **upvote** the best ones at `/community`,
  so real players help grow and improve the question bank. A lightweight
  built-in "not a robot" check (a signed arithmetic captcha) keeps out drive-by
  spam. This is the one part of the app that is **persisted** (see below).
- **Play official or community questions** — in the lobby the quiz master picks
  the question **source** (📚 Official bank or 🗳️ Community) right next to the
  language selector. Pick *Community* and the game plays the **top-voted**
  submissions for the chosen language & difficulty, so the crowd's votes
  directly shape what gets played (falls back to the official bank if the
  community pool is empty).
- **In-memory gameplay** — live quizzes have no database, no accounts, no stored
  data; everything lives in RAM and is wiped when the server restarts. Security
  is intentionally out of scope (family use, no sensitive data).

## Run locally

```bash
go run .
# then open http://localhost:8080
```

Set a custom port with `PORT=9000 go run .`.

To exercise the whole flow on one machine, open the admin link in one browser
and the join link (or `/j/<CODE>`) in another (or a phone on the same network).

## How it works

| Page | Who | Purpose |
|------|-----|---------|
| `/` | anyone | Pick a difficulty and create a quiz |
| `/admin/{id}?t={token}` | quiz master | Scoreboard, QR code, run the quiz |
| `/j/{id}` | attendees | Join as a team (scan the QR code) |
| `/play/{id}` | team captains | Answer questions live |
| `/community` | anyone | Browse & upvote community-submitted questions |
| `/community/new` | anyone | Submit a question (captcha-gated) |

State is polled over a small JSON API every ~1.5s, so no WebSockets are needed.
The UI is styled with [Tailwind CSS](https://tailwindcss.com) (Play CDN) and QR
codes are generated server-side with
[`go-qrcode`](https://github.com/skip2/go-qrcode). Templates are embedded into
the binary via `go:embed`, so the compiled binary is fully self-contained.

## Deploy to Railway

This repo is ready for [Railway](https://railway.app):

1. Create a new Railway project → **Deploy from GitHub repo** → pick this repo.
2. Railway detects the `Dockerfile` and builds it (config in `railway.json`).
3. Railway injects a `PORT` env var, which the server reads automatically.
4. A healthcheck is configured at `/healthz`.

For **community questions to survive restarts/redeploys**, give the service a
persistent store. The app picks one automatically:

| Env var | Effect |
|---------|--------|
| `DATABASE_URL` | If set, community questions are stored in **Postgres** (add a Postgres database in Railway and reference its `DATABASE_URL`). Recommended for production. |
| `DATA_DIR` | Otherwise questions are kept in a JSON file at `$DATA_DIR/community.json` (default `./data`). Point this at a mounted **Railway Volume** to persist. |
| `CAPTCHA_SECRET` | Optional. Signs captcha tokens; if unset a random secret is generated at startup (a restart simply invalidates in-flight challenges). |

Without any of these the quiz still runs perfectly; only the community board's
data is ephemeral. Once deployed, open the public URL and start a quiz.

## Updating the questions

All questions live in [`questions.go`](questions.go), grouped by difficulty.
Each is multiple-choice with the correct option index and a fun fact shown on
reveal. Edit, rebuild, redeploy.
