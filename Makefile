.PHONY: all test

all: test build_receiver

build_receiver:
	go build -o bin/receiver ./cli/receiver

test:
	go test ./...