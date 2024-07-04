
PARENT_DIR:=$(shell dirname $(CURDIR))
LOCAL_BIN:=$(PARENT_DIR)/bin

.PHONY: generate
generate:
	$(info Generate Note Service stubs...)

	PATH="$(LOCAL_BIN):$(PATH)" oapi-codegen -config ./build/note.oapi.yaml ./api/v1/note.openapi.yaml
	PATH="$(LOCAL_BIN):$(PATH)" oapi-codegen -config ./build/card.oapi.yaml ./api/v1/card.openapi.yaml
	go mod tidy
