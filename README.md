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
make build
./gbot start
```

When the bot starts you will see:

```
gBot <version> started
```

The version is taken from git automatically: `<last tag>-<short commit>`, e.g. `v0.0.0-ab75275`.

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

| Target              | Description                                                                                       |
|---------------------|---------------------------------------------------------------------------------------------------|
| `make run C=<args>` | Run the app with `go run`. `C` is passed as CLI arguments, e.g. `make run C=start`                |
| `make format`       | Format all Go sources with `gofmt -s -w`                                                          |
| `make lint`         | Run `go vet ./...`                                                                                |
| `make test`         | Run `go test -v ./...`                                                                            |
| `make get`          | Download Go module dependencies                                                                   |
| `make build`        | Build the `gbot` binary for the host platform (or `TARGETOS`/`TARGETARCH` if set)                 |
| `make linux`        | Build for Linux amd64 → `gbot`                                                                    |
| `make arm`          | Build for Linux arm64 → `gbot`                                                                    |
| `make macos`        | Build for macOS arm64 (Apple Silicon) → `gbot`                                                    |
| `make windows`      | Build for Windows amd64 → `gbot.exe`                                                              |
| `make image`        | Build a Docker image for the host architecture                                                    |
| `make push`         | Push the image to the registry (`ghcr.io/vystepanenko` by default)                                |
| `make clean`        | Remove built binaries and the Docker image (`docker rmi <IMAGE_TAG>`)                             |

## Multi-platform builds and testing (for QA)

This section describes how to build the bot and run tests on different platforms and architectures:
Linux, arm, macOS and Windows.

### Requirements

- `make`, `git`, Go 1.25+ (for local binaries)
- Docker 23+ with BuildKit enabled (the default). A plain `docker build` is used, **no buildx and no QEMU** needed
- **Windows:** run `make` from Git Bash, MSYS2 or WSL (`uname` is needed to detect the host platform); `cmd.exe` and PowerShell are not supported
- Push to the registry only: `docker login ghcr.io`

### Building binaries

Each target cross-compiles a static binary (`CGO_ENABLED=0`) into the project root:

```bash
make linux     # Linux   amd64 → ./gbot
make arm       # Linux   arm64 → ./gbot
make macos     # macOS   arm64 → ./gbot
make windows   # Windows amd64 → ./gbot.exe
```

Any other combination can be built directly:

```bash
make build TARGETOS=darwin TARGETARCH=amd64   # macOS on Intel
make build TARGETOS=linux  TARGETARCH=arm64
```

> All targets except `windows` write to the same `./gbot` file, so the last build overwrites the previous one.

Check the result:

```bash
file gbot        # e.g. "ELF 64-bit LSB executable, ARM aarch64" or "Mach-O 64-bit arm64 executable"
./gbot version   # only on a matching host
```

### Docker image for test runs

`make image` builds an image **for the architecture of the host where it is run**:

| Host                         | Image platform | Emulation      |
|------------------------------|----------------|----------------|
| Linux x86_64                 | `linux/amd64`  | none           |
| Linux arm64 (Raspberry Pi, Graviton) | `linux/arm64` | none   |
| macOS Apple Silicon (M1+)    | `linux/arm64`  | none (no Rosetta, no QEMU) |
| macOS Intel                  | `linux/amd64`  | none           |
| Windows x86_64 (Docker Desktop) | `linux/amd64` | none          |

How this works without emulation: the builder stage runs on the native platform of the build host
(`FROM --platform=$BUILDPLATFORM`), and Go cross-compiles the binary for the target platform
(`GOOS=$TARGETOS GOARCH=$TARGETARCH`). The final image is `scratch` and contains only the binary and CA certificates.

Images are tagged as:

```
ghcr.io/vystepanenko/gbot:<version>-linux-<arch>
# e.g. ghcr.io/vystepanenko/gbot:v0.0.0-ab75275-linux-arm64
```

The base image (`quay.io/projectquay/golang`) and the registry (`ghcr.io`) are not on Docker Hub,
so Docker Hub licensing and pull rate limits do not apply.

### Typical test flow

```bash
# 1. Build the image for this host
make image

# 2. Find the tag
docker image ls ghcr.io/vystepanenko/gbot

# 3. Run the binary in the container
docker run --rm ghcr.io/vystepanenko/gbot:<tag> version
docker run --rm -e TELE_TOKEN ghcr.io/vystepanenko/gbot:<tag> start

# 4. Clean up: removes the binaries and the image built in step 1
make clean
```

An image for a different architecture can also be built (the binary is still cross-compiled, no emulation during the build).
Running it, however, needs a host with that architecture:

```bash
make image TARGETARCH=arm64
make clean TARGETARCH=arm64
```

> `make clean` removes the image whose tag is computed from the **current** commit.
> Run it before committing or switching branches, otherwise remove the old image manually with `docker rmi <tag>`.

## Project structure

```
.
├── main.go               # entry point, calls cmd.Execute()
├── Makefile              # build, cross-compile, image and clean targets
├── Dockerfile            # multi-stage build: quay.io golang builder → scratch
├── .dockerignore         # files excluded from the Docker build context
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
