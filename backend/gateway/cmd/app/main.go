package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/knowdateapp/knowdateapp/backend/gateway/internal/pb/knowledge/api/v1/collection"
	"github.com/knowdateapp/knowdateapp/backend/gateway/internal/pb/knowledge/api/v1/knowledge_base"
	"github.com/knowdateapp/knowdateapp/backend/gateway/internal/pb/knowledge/api/v1/topic"
)

func main() {
	const (
		ServerAddressHTTP = ":80"
	)

	var (
		knowledgeServiceAddressGRPC = flag.String("knowledge-service-address", ":8012", "Knowledge Base service address")
		userServiceAddressGRPC      = flag.String("user-service-address", ":8022", "User service address")
		_                           = userServiceAddressGRPC
	)

	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	options := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	services := []RegisterFunc{
		CaptureRegisterFuncWithEndpoint(*knowledgeServiceAddressGRPC, knowledge_base.RegisterKnowledgeBaseServiceHandlerFromEndpoint),
		CaptureRegisterFuncWithEndpoint(*knowledgeServiceAddressGRPC, collection.RegisterCollectionServiceHandlerFromEndpoint),
		CaptureRegisterFuncWithEndpoint(*knowledgeServiceAddressGRPC, topic.RegisterTopicServiceHandlerFromEndpoint),
	}

	for _, registerFunc := range services {
		err := registerFunc(ctx, mux, options)
		if err != nil {
			log.Fatalln("Can't start application:", err)
		}
	}

	// Start application.
	log.Println("Gateway server startup on port:", ServerAddressHTTP)

	err := http.ListenAndServe(ServerAddressHTTP, mux)
	if err != nil {
		log.Fatalln("Can't start application:", err)
	}
}

type RegisterFuncWithEndpoint = func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error
type RegisterFunc = func(context.Context, *runtime.ServeMux, []grpc.DialOption) error

func CaptureRegisterFuncWithEndpoint(endpoint string, register RegisterFuncWithEndpoint) RegisterFunc {
	return func(ctx context.Context, mux *runtime.ServeMux, options []grpc.DialOption) error {
		return register(ctx, mux, endpoint, options)
	}
}
