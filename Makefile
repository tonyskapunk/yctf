.PHONY: lint
lint: golangci-lint
	$(GOLANGCI_LINT) run

# install golangci-lint
GOLANGCI_LINT = $(shell pwd)/bin/golangci-lint 	# to this path
GOLANGCI_LINT_VERSION ?= v1.49.0 				# at this version
golangci-lint: $(GOLANGCI_LINT)
$(GOLANGCI_LINT):
	$(call go-install-tool,$(GOLANGCI_LINT),github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION))

# go-install-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-install-tool
@[ -f $(1) ] || { \
GOBIN=$(PROJECT_DIR)/bin go install $(2) ;\
}
endef