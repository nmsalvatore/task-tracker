.DEFAULT_GOAL := build

.PHONY:fmt vet build clean

fmt:
	go fmt ./...

vet: fmt
	go fmt ./...

build: vet
	go build

clean:
	go clean
