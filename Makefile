BINARY_NAME=pay-friends-back.exe

.PHONY: all build run qrun test clean

all: build test

build:
	go build -o ${BINARY_NAME} .

run:
	./${BINARY_NAME} -env=dev

qrun:
	go run . -env=dev

test:
	go test ./... -v

clean:
	go clean
	rm -rf ${BINARY_NAME}