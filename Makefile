BINARY_NAME=bin/pay-friends-back

.PHONY: all build run qrun test clean

all: build test

build:
	go build -o ${BINARY_NAME} .

run:
	./${BINARY_NAME}

qrun:
	go run .

test:
	go test ./... -v

clean:
	go clean
	rm -rf ${BINARY_NAME}