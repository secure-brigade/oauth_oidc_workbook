PROJECT_PATH=$(shell pwd)
export GO_PATH=${GOPATH}

.PHONY: help
all: help

# load .env
export ENV ?= ${ENV}
ifneq (,$(wildcard ./.env))
	include .env
endif

help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo

## mod-download: Download modules from go.sum
mod-download:
	@go mod download

## mod-tidy: Tidy up modules
mod-tidy:
	@go mod tidy

# migrate-up: migrate tables
migrate-up: gen-model-from-schema migrate-schema
gen-model-from-schema:
	@go generate ./ent
migrate-schema:
	@go run migration/main.go

## $(GO_PATH)/bin/air: download tool for live reloading
$(GO_PATH)/bin/air:
	@go get -u github.com/cosmtrek/air

## serve-locally: Run server in local
serve-locally: $(GO_PATH)/bin/air
	@air #