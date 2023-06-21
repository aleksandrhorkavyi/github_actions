DIR = $(shell pwd)
BIN = $(DIR)/.bin
$(shell [ -f .bin ] || mkdir -p $(BIN))
PATH := $(BIN):$(PATH)

GOLANGCI_LINT = $(BIN)/golangci-lint

.PHONY: .install-linter
.install-linter:
	### INSTALL GOLANGCI-LINT ###
	[ -f $(BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(BIN) v1.53.3

.PHONY: lint
lint: .install-linter
	### RUN GOLANGCI-LINT ###
	$(GOLANGCI_LINT) run ./... --config=./build/ci/golangci.yml
