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
- **In-memory only** — no database, no accounts, no stored data. Everything
  lives in RAM and is wiped when the server restarts. Security is intentionally
  out of scope (family use, no sensitive data).

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

No environment variables are required. Once deployed, open the public URL and
start a quiz.

## Updating the questions

All questions live in [`questions.go`](questions.go), grouped by difficulty.
Each is multiple-choice with the correct option index and a fun fact shown on
reveal. Edit, rebuild, redeploy.
