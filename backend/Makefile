
LOCAL_BIN:=$(CURDIR)/bin
VENDOR_PROTO:=$(CURDIR)/vendor.proto

.PHONY: install-bin-deps
install-bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.11.2

.PHONY: proto-deps-vendor
proto-deps-vendor:
	$(info Vendoring...)

	mkdir -p $(VENDOR_PROTO)
	mkdir -p $(VENDOR_PROTO)/google/api
	mkdir -p $(VENDOR_PROTO)/google/protobuf
	mkdir -p $(VENDOR_PROTO)/protoc-gen-openapiv2/options

	curl -so $(VENDOR_PROTO)/google/api/annotations.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto
	curl -so $(VENDOR_PROTO)/google/api/field_behavior.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/field_behavior.proto
	curl -so $(VENDOR_PROTO)/google/api/http.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto
	curl -so $(VENDOR_PROTO)/google/api/httpbody.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/httpbody.proto
	curl -so $(VENDOR_PROTO)/google/api/field_behavior.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/field_behavior.proto
	curl -so $(VENDOR_PROTO)/google/protobuf/descriptor.proto https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/descriptor.proto
	curl -so $(VENDOR_PROTO)/google/protobuf/timestamp.proto https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/timestamp.proto
	curl -so $(VENDOR_PROTO)/protoc-gen-openapiv2/options/annotations.proto https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/main/protoc-gen-openapiv2/options/annotations.proto
	curl -so $(VENDOR_PROTO)/protoc-gen-openapiv2/options/openapiv2.proto https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/main/protoc-gen-openapiv2/options/openapiv2.proto
