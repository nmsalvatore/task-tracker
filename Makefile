.DEFAULT_GOAL := build

.PHONY:fmt vet build clean

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o task ./cmd/task

clean:
	go clean
	rm -f task

test:
	go test ./cmd/task
