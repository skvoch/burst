.PH0NY: build

build:
	go build -o ./bin/ -v ./cmd/burst
	go build -o ./bin/ -v ./cmd/telegram

test:
	go test -v -race ./...

.DEFAULT_GOAL := build
