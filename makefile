SHELL := /bin/bash
MAKEFLAGS += --warn-undefined-variables
.DEFAULT_GOAL := build
.PHONY: *

tools: ## Download and install all dev/code tools
	@echo "==> Installing dev tools"
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

build:
	@echo "==> Building go-centreon"
	go build -o /dev/null .

check:
	@echo "==> Checking go-centreon"
	gometalinter \
    		--deadline 10m \
    		--vendor \
    		--sort="path" \
    		--aggregate \
    		--enable-gc \
    		--disable-all \
    		--enable goimports \
    		--enable misspell \
    		--enable vet \
    		--enable deadcode \
    		--enable varcheck \
    		--enable ineffassign \
    		--enable gofmt \
				./...
