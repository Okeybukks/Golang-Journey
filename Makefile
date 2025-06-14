.DEFAULT_GOAL := build

.PHONY: format validate build

format:
	go fmt ./...

validate: format
	go vet ./...

build: validate
	go build -o todo

clean:
	go clean ./...