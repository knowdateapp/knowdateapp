
PARENT_DIR:=$(shell dirname $(CURDIR))
LOCAL_BIN:=$(PARENT_DIR)/bin
VENDOR_PROTO:=$(PARENT_DIR)/vendor.proto

.PHONY: generate
generate:
	$(info Generate User Service stubs...)
	echo $(VENDOR_PROTO)

	PATH="$(LOCAL_BIN):$(PATH)" protoc -I $(VENDOR_PROTO) -I . \
 			--go_out ./internal/pb --go_opt paths=source_relative \
     		--go-grpc_out ./internal/pb --go-grpc_opt paths=source_relative \
     		api/v1/auth/auth.proto \
			api/v1/user/user.proto && \
		go mod tidy
