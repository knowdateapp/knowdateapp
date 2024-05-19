
PARENT_DIR:=$(shell dirname $(CURDIR))
LOCAL_BIN:=$(PARENT_DIR)/bin

.PHONY: generate
generate:
	$(info Generate User Service stubs...)

	PATH="$(LOCAL_BIN):$(PATH)" oapi-codegen -config ./build/user.oapi.yaml ./api/v1/user.openapi.yaml
	PATH="$(LOCAL_BIN):$(PATH)" oapi-codegen -config ./build/auth.oapi.yaml ./api/v1/auth.openapi.yaml
	go mod tidy
