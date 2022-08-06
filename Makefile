.PHONY: all test
TEST_CONFIG_PATH = $(shell pwd)/configs/config.test.yaml

all: build_receiver build_packet_gen

docker:
	docker build -t egts:latest .

build_receiver:
	go build -o bin/receiver ./cli/receiver

build_packet_gen:
	go build -o bin/packet_gen ./cli/packet-gen

test:
	docker-compose -f docker-compose-test-env.yml up -d
	sleep 10
	TEST_CONFIG=$(TEST_CONFIG_PATH) go test ./...
	make clean

clean:
	go clean -testcache
	docker-compose -f docker-compose-test-env.yml down
