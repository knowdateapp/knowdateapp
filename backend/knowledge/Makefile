
PARENT_DIR:=$(shell dirname $(CURDIR))
LOCAL_BIN:=$(PARENT_DIR)/bin
VENDOR_PROTO:=$(PARENT_DIR)/vendor.proto

.PHONY: generate
generate:
	$(info Generate Knowledge Base Service stubs...)
	echo $(VENDOR_PROTO)

	PATH="$(LOCAL_BIN):$(PATH)" protoc -I $(VENDOR_PROTO) -I . \
     			--go_out ./internal/pb --go_opt paths=source_relative \
         		api/v1/common/topic.proto \
         		api/v1/common/collection.proto \
         		api/v1/common/knowledge_base.proto

	PATH="$(LOCAL_BIN):$(PATH)" protoc -I $(VENDOR_PROTO) -I . \
 			--go_out ./internal/pb --go_opt paths=source_relative \
     		--go-grpc_out ./internal/pb --go-grpc_opt paths=source_relative \
     		api/v1/topic/topic.proto \
     		api/v1/collection/collection.proto \
			api/v1/knowledge_base/knowledge_base.proto && \
		go mod tidy
