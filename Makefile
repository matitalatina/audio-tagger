.PHONY: install build test

build:
	@echo "Start building"
	@go build -o build/audio-tagger

install:
	@echo "Getting dependencies"
	go get -t -v ./...

test:
	go test

