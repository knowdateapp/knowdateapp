package knowledge_base

import (
	"google.golang.org/grpc"

	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/knowledge_base"
)

type Implementation struct {
	desc.UnimplementedKnowledgeBaseServiceServer
}

func NewKnowledgeBaseService() *Implementation {
	return &Implementation{}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.KnowledgeBaseServiceServer) {
	desc.RegisterKnowledgeBaseServiceServer(s, srv)
}
