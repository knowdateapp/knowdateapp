package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/app/v1/collection"
	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/app/v1/knowledge_base"
	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/app/v1/topic"
)

func main() {
	const (
		ServerAddressGRPC = ":82"
	)

	// GRPC server setup.
	server := grpc.NewServer()

	knowledge_base.RegisterGRPCServer(server, knowledge_base.NewKnowledgeBaseService())
	collection.RegisterGRPCServer(server, collection.NewCollectionService())
	topic.RegisterGRPCServer(server, topic.NewTopicService())

	reflection.Register(server)

	// Start application.
	log.Println("Knowledge server startup on port:", ServerAddressGRPC)

	lis, err := net.Listen("tcp", ServerAddressGRPC)
	if err != nil {
		log.Fatalln("Can't start application:", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Fatalln("Can't start application:", err)
	}
}
