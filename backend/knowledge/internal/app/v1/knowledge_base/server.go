package knowledge_base

import (
	"google.golang.org/grpc"

	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/domain/services"
	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/knowledge_base"
)

type Implementation struct {
	desc.UnimplementedKnowledgeBaseServiceServer

	knowledgeBaseService *services.KnowledgeBaseService
}

func NewKnowledgeBaseService(knowledgeBaseService *services.KnowledgeBaseService) *Implementation {
	return &Implementation{
		knowledgeBaseService: knowledgeBaseService,
	}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.KnowledgeBaseServiceServer) {
	desc.RegisterKnowledgeBaseServiceServer(s, srv)
}
