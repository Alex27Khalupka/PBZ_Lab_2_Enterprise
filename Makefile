.PHONY: build
build:
	go build -v ./cmd/apiserver
.PHONY: export_variables
.DEFAULT_GOAL := build