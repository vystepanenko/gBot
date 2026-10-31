APP      := $(shell basename $(shell git remote get-url origin) .git | tr '[:upper:]' '[:lower:]')
REGISTRY ?= ghcr.io/vystepanenko
VERSION  := $(shell git describe --tags --abbrev=0 2>/dev/null || echo v0.0.0)-$(shell git rev-parse --short HEAD)
LDFLAGS  := -X github.com/vystepanenko/gBot/cmd.appVersion=$(VERSION)

# Platform of the host where make is executed
HOST_OS   := $(shell uname -s | tr '[:upper:]' '[:lower:]' | sed -E 's/^(mingw|msys|cygwin).*/windows/')
HOST_ARCH := $(shell uname -m | sed -e 's/x86_64/amd64/' -e 's/aarch64/arm64/')

# Target platform (host by default, override: make build TARGETOS=linux TARGETARCH=arm64)
TARGETOS   ?= $(HOST_OS)
TARGETARCH ?= $(HOST_ARCH)

BINARY    := $(APP)
IMAGE_TAG  = $(REGISTRY)/$(APP):$(VERSION)-linux-$(TARGETARCH)

.PHONY: run format lint test get build linux arm macos windows image push clean

run:
	go run main.go $(C)

format:
	gofmt -s -w ./

lint:
	go vet ./...

test:
	go test -v ./...

get:
	go mod download

build: format get
	CGO_ENABLED=0 GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build -v -o $(BINARY)$(EXT) -ldflags "$(LDFLAGS)"

linux:
	$(MAKE) build TARGETOS=linux TARGETARCH=amd64

arm:
	$(MAKE) build TARGETOS=linux TARGETARCH=arm64

macos:
	$(MAKE) build TARGETOS=darwin TARGETARCH=arm64

windows:
	$(MAKE) build TARGETOS=windows TARGETARCH=amd64 EXT=.exe

# Container image for the host architecture (Linux container, native arch: no QEMU / Rosetta)
image:
	docker build . --platform linux/$(TARGETARCH) --build-arg VERSION=$(VERSION) -t $(IMAGE_TAG)

push:
	docker push $(IMAGE_TAG)

clean:
	rm -f $(BINARY) $(BINARY).exe
	-docker rmi $(IMAGE_TAG)
