
SHELL := /bin/bash
GO := GO111MODULE=on GO15VENDOREXPERIMENT=1 go
GO_NOMOD := GO111MODULE=off go
GOTEST := $(GO) test
PACKAGE_DIRS := $(shell $(GO) list ./... | grep -v /vendor/)
GO_DEPENDENCIES := $(shell find . -type f -name '*.go')

build: 
	$(GO) build ./...

linux: build

test: build
	$(GOTEST) -coverprofile=coverage.out ./...

test1: ## Runs single test specified by test name and optional package, eg 'make test1 TEST=TestGitCLI'
	$(GOTEST) -v ./pkg/log -run $(TEST)

get-fmt-deps: ## Install test dependencies
	$(GO_NOMOD) get golang.org/x/tools/cmd/goimports

fmt: importfmt ## Format the code
	@FORMATTED=`$(GO) fmt $(PACKAGE_DIRS)`
	@([[ ! -z "$(FORMATTED)" ]] && printf "Fixed unformatted files:\n$(FORMATTED)") || true

importfmt: get-fmt-deps
	@echo "Formatting the imports..."
	goimports -w $(GO_DEPENDENCIES)

.PHONY: lint
lint: ## Lint the code
	./hack/linter.sh

.PHONY: modtidy
modtidy:
	$(GO) mod tidy

.PHONY: coverage
coverage:
	$(GO) tool cover -html=coverage.out

.PHONY: cover
cover:
	$(GO) tool cover -func coverage.out | grep total

.PHONY: code-generate
code-generate:
	./hack/generate.sh

.PHONY: docs gen-schema
docs: generate-refdocs

.PHONY: gen-schema
gen-schema:
	mkdir -p schema
	go run cmd/schemagen/main.go

include Makefile.codegen

