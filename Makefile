.PHONY: all test

all: build_receiver build_packet_gen

build_receiver:
	go build -o bin/receiver ./cli/receiver

build_packet_gen:
	go build -o bin/packet_gen ./cli/packet-gen

test:
	go test ./...