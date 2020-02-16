.PH0NY: build

build:
	go build -v ./cmd/burst

test:
	go test -v -race ./...

.DEFAULT_GOAL := build