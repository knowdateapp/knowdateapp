
PARENT_DIR:=$(shell dirname $(CURDIR))
LOCAL_BIN:=$(PARENT_DIR)/bin
VENDOR_PROTO:=$(PARENT_DIR)/vendor.proto
VENDOR_PROTO_LOCAL:=$(CURDIR)/vendor.proto

.PHONY: generate
generate:
	$(info Generate Gateway Service stubs...)
	$(info Vendor proto $(VENDOR_PROTO))

	cat dependencies.txt | python3 ./build/vendor-proto.py

	# Generate HTTP to GRPC proxy implementation for Knowledge Service.
	PATH="$(LOCAL_BIN):$(PATH)" protoc -I "$(VENDOR_PROTO)" -I "$(VENDOR_PROTO_LOCAL)/knowledge" \
			--go_out ./internal/pb/knowledge --go_opt paths=source_relative \
			api/v1/common/topic.proto \
			api/v1/common/collection.proto \
			api/v1/common/knowledge_base.proto

	PATH="$(LOCAL_BIN):$(PATH)" protoc -I "$(VENDOR_PROTO)" -I "$(VENDOR_PROTO_LOCAL)/knowledge" \
			--go_out ./internal/pb/knowledge --go_opt paths=source_relative \
     		--go-grpc_out ./internal/pb/knowledge --go-grpc_opt paths=source_relative \
     		--grpc-gateway_out ./internal/pb/knowledge --grpc-gateway_opt paths=source_relative \
     		--openapiv2_out ./internal/pb/knowledge --openapiv2_opt output_format=yaml --openapiv2_opt json_names_for_fields=false \
     		api/v1/topic/topic.proto \
			api/v1/collection/collection.proto \
			api/v1/knowledge_base/knowledge_base.proto && \
		go mod tidy

	# Generate HTTP to GRPC proxy implementation for User Service.
	PATH="$(LOCAL_BIN):$(PATH)" protoc -I "$(VENDOR_PROTO)" -I "$(VENDOR_PROTO_LOCAL)/user" \
			--go_out ./internal/pb/user --go_opt paths=source_relative \
			--go-grpc_out ./internal/pb/user --go-grpc_opt paths=source_relative \
			--grpc-gateway_out ./internal/pb/user --grpc-gateway_opt paths=source_relative \
			--openapiv2_out ./internal/pb/user --openapiv2_opt output_format=yaml --openapiv2_opt json_names_for_fields=false \
			api/v1/auth/auth.proto \
			api/v1/user/user.proto && \
		go mod tidy
