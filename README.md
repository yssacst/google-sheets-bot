# Google Sheets Bot

A lightweight Go bot that reads a Google Sheets spreadsheet every day, checks whether a user is scheduled for the next day, and sends a notification through a configurable notification provider.

The project was designed to be simple, maintainable and extensible, following a lightweight Hexagonal Architecture (Ports & Adapters).

## Features

- Read data from Google Sheets
- Check if a user is on duty for the next day
- Send notifications through an abstract notification interface
- Current notification implementation: ntfy.sh
- Daily execution using GitHub Actions
- Configuration through environment variables
- Structured logging
- Lightweight architecture with separation of concerns

---

## Architecture

```
                +----------------+
                | GitHub Actions |
                +--------+-------+
                         |
                         v
                  +-------------+
                  |   main.go   |
                  +------+------+ 
                         |
                         v
                  +-------------+
                  |     App     |
                  | Orchestrator|
                  +------+------+ 
                         |
        +----------------+----------------+
        |                                 |
        v                                 v
+---------------+               +------------------+
| Google Sheets |               | Notification     |
| Adapter       |               | Adapter          |
+-------+-------+               +--------+---------+
        |                                |
        v                                v
 Google Sheets API                 ntfy.sh (today)
                            Telegram (future)
                            Slack (future)
                            Discord (future)
```

The business rules never know **how** the notification is delivered.

Only the adapter knows that.

---

## Project Structure

```
.
├── cmd/
│   └── bot/
│       └── main.go
│
├── internal/
│   ├── app/
│   ├── config/
│   ├── core/
│   ├── logger/
│   ├── notifier/
│   └── sheets/
│
├── docs/
├── .github/
│   └── workflows/
│       └── bot.yml
│
├── .env.example
├── go.mod
└── README.md
```

---

## Requirements

- Go 1.24+
- Google Cloud Project
- Google Sheets API enabled
- Service Account
- Google Spreadsheet
- GitHub Actions (optional)

---

## Configuration

Create a `.env` file.

Example:

```env
LOG_LEVEL=INFO

SPREADSHEET_ID=
SHEET_NAME=

USER_NAME=

API_URL=

GOOGLE_CREDENTIALS=
```

See:

```
docs/howto.txt
```

for the complete Google Cloud setup.

---

## Running locally

Clone the project

```bash
git clone https://github.com/yssacst/google-sheets-bot.git

cd google-sheets-bot
```

Install dependencies

```bash
go mod download
```

Run

```bash
go run ./cmd/bot
```

---

## GitHub Actions

The bot can execute automatically every day.

Example:

```yaml
schedule:
  - cron: "20 22 * * *"
```

22:20 UTC = 19:20 (America/São_Paulo)

---

## Notification Providers

The notification system is abstracted behind an interface.

Current implementation:

- ✅ ntfy.sh

Future implementations can include:

- Telegram
- Discord
- Slack
- Email
- Microsoft Teams

without changing the business rules.

---

## Configuration Files

| File | Description |
|-------|-------------|
| `.env.example` | Environment variables example |
| `docs/howto.txt` | Google Cloud configuration |
| `.github/workflows/bot.yml` | GitHub Actions workflow |

---

## Development

Run tests

```bash
go test ./...
```

Format code

```bash
go fmt ./...
```

Run vet

```bash
go vet ./...
```

---

## Versioning

This project follows Semantic Versioning.

Examples:

```
v1.0.0
v1.0.1
v1.1.0
v2.0.0
```

Versions are injected during build using Go ldflags.

---

## Roadmap

- [x] Google Sheets integration
- [x] ntfy notification
- [x] GitHub Actions
- [ ] Telegram notifier
- [ ] Slack notifier
- [ ] Multiple notification providers
- [ ] Unit tests
- [ ] Integration tests
- [ ] Docker support

---

## License

MIT
