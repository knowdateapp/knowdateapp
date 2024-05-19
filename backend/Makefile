
LOCAL_BIN:=$(CURDIR)/bin

.PHONY: generate
generate: install-bin-deps .generate

.PHONY: fast-generate
fast-generate: .generate

.PHONY: .generate
.generate:
	$(info Generate OpenAPI stubs...)

	$(MAKE) -C user generate
	$(MAKE) -C note generate

.PHONY: install-bin-deps
install-bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(LOCAL_BIN) go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0
