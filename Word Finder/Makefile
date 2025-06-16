.DEFAULT_GOAL := run

.PHONY : format validate build 
format:
	go fmt ./...

validate: format
	go vet ./...

run: validate
	go run .
build: validate
	go build -o finder

clean:
	go clean ./...