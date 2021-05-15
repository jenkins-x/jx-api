
SHELL := /bin/bash
GO := GO111MODULE=on GO15VENDOREXPERIMENT=1 go
GO_NOMOD := GO111MODULE=off go
GOTEST := $(GO) test
PACKAGE_DIRS := $(shell $(GO) list ./... | grep -v /vendor/)
GO_DEPENDENCIES := $(shell find . -type f -name '*.go')

.PHONY: build
build: 
	$(GO) build ./...

.PHONY: linux
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

# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true,preserveUnknownFields=false"

manifests: controller-gen ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

CONTROLLER_GEN = $(shell pwd)/bin/controller-gen
controller-gen: ## Download controller-gen locally if necessary.
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.1)


include Makefile.codegen


# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go get $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef

