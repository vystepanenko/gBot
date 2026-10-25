# gBot

A simple Telegram bot written in Go as a test task for the **DevOps and Kubernetes** course.

Bot: [t.me/devops_g_bot](https://t.me/devops_g_bot)

Built with [cobra](https://github.com/spf13/cobra) for the CLI and [telebot v4](https://gopkg.in/telebot.v4) for the Telegram Bot API.

## Requirements

- Go 1.25+
- `make` (optional, for the shortcuts below)
- A Telegram bot token from [@BotFather](https://t.me/BotFather)

## Getting a bot token

1. Open [@BotFather](https://t.me/BotFather) in Telegram.
2. Send `/newbot` and follow the instructions (choose a name and a username ending in `bot`).
3. Copy the token BotFather gives you. It looks like `123456789:AAH...`.

Keep the token secret: anyone who has it can control your bot.

## Configuration

The bot reads its token from the `TELE_TOKEN` environment variable:

```bash
read -s TELE_TOKEN   # paste the token, it won't be echoed
export TELE_TOKEN
```

## Running

### With make

```bash
make run C=start
```

### With go

```bash
go run main.go start
```

### From a built binary

```bash
make build C=v1.0.0
./gBot start
```

When the bot starts you will see:

```
gBot v1.0.0 started
```

Now open your bot in Telegram and send `/start`.

> Only one instance of the bot can run with the same token at a time. A second one fails with `409 Conflict`.

## CLI commands

| Command         | Description                                   |
|-----------------|-----------------------------------------------|
| `gBot start`    | Start the Telegram bot (alias for `gBot gbot`) |
| `gBot version`  | Print the application version                 |
| `gBot help`     | Show CLI help                                 |

## Bot commands

| Command         | Description                                              |
|-----------------|----------------------------------------------------------|
| `/start`        | Start the bot                                            |
| `/help`         | Show the list of available commands                      |
| `/version`      | Show the bot version                                     |
| `/conversation` | Start a short conversation with inline **Yes/No** buttons |
| `/about`        | Information about the bot                                |

The bot also replies to plain text messages:

| Message          | Reply                                        |
|------------------|----------------------------------------------|
| `hello` / `hi`   | Greets you by your first name                |
| `version`        | Shows the bot version                        |
| anything else    | Suggests typing `/help`                      |

## Make targets

| Target                 | Description                                                                                         |
|------------------------|-----------------------------------------------------------------------------------------------------|
| `make run C=<args>`    | Run the app with `go run`. `C` is passed as CLI arguments, e.g. `make run C=start` or `make run C=version` |
| `make build C=<ver>`   | Build the `gBot` binary. `C` is embedded as the app version via `-ldflags -X`, e.g. `make build C=v1.0.0` |
| `make format`          | Format all Go sources with `gofmt -s -w`                                                            |

If `C` is omitted in `make build`, the version will be empty.

## Project structure

```
.
├── main.go               # entry point, calls cmd.Execute()
├── cmd/                  # CLI commands (cobra)
│   ├── root.go           # root command
│   ├── gbot.go           # `gbot` (`start`) command: creates and starts the bot
│   └── version.go        # `version` command and appVersion variable
└── internal/bot/         # Telegram bot logic
    ├── bot.go            # command list and handler registration (Register)
    ├── start.go          # /start
    ├── help.go           # /help (generated from the command list)
    ├── text.go           # plain text messages (OnText)
    ├── version.go        # /version
    ├── about.go          # /about
    └── conversation.go   # /conversation and inline buttons
```

## Adding a new bot command

1. Declare it in `internal/bot/bot.go`:
   ```go
   cmdPing = command{name: "/ping", description: "check that the bot is alive"}
   ```
2. Add it to the `commands` slice so it appears in `/help`.
3. Create a handler in its own file, e.g. `internal/bot/ping.go`:
   ```go
   func handlePing(ctx telebot.Context) error {
   	return ctx.Send("pong")
   }
   ```
4. Register it in `Register`:
   ```go
   gbot.Handle(cmdPing.name, handlePing)
   ```
