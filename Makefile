.PHONY: build
build:
	go build -v ./cmd/apiserver
.PHONE: test
test:
	go test ./pkg/service
.DEFAULT_GOAL := build