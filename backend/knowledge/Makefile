
PARENT_DIR:=$(shell dirname $(CURDIR))
LOCAL_BIN:=$(PARENT_DIR)/bin
VENDOR_PROTO:=$(PARENT_DIR)/vendor.proto

.PHONY: generate
generate:
	$(info Generate Knowledge Base Service stubs...)
	echo $(VENDOR_PROTO)

	PATH="$(LOCAL_BIN):$(PATH)" protoc -I $(VENDOR_PROTO) -I . \
     			--go_out ./internal/pb --go_opt paths=source_relative \
         		api/v1/common/page.proto \
         		api/v1/common/collection.proto \
         		api/v1/common/base.proto && \
    		go mod tidy

	PATH="$(LOCAL_BIN):$(PATH)" protoc -I $(VENDOR_PROTO) -I . \
 			--go_out ./internal/pb --go_opt paths=source_relative \
     		--go-grpc_out ./internal/pb --go-grpc_opt paths=source_relative \
     		--grpc-gateway_out ./internal/pb --grpc-gateway_opt paths=source_relative \
     		--openapiv2_out ./internal/pb --openapiv2_opt output_format=yaml --openapiv2_opt json_names_for_fields=false \
     		api/v1/page/page.proto \
     		api/v1/collection/collection.proto \
			api/v1/base/base.proto && \
		go mod tidy
