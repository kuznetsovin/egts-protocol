.PHONY: all test

all: test build

build:
	go build -o bin/receiver ./app 

test:
	go test ./...