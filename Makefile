run:
	go run main.go $(C)

format:
	gofmt -s -w ./

build:
	go build -ldflags "-X="github.com/vystepanenko/gBot/cmd.appVersion=$(C)
