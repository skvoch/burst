.PH0NY: build

build:
	go build -v ./cmd/burst
	go build -v ./cmd/telegram

test:
	go test -v -race ./...

.DEFAULT_GOAL := build