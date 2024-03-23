package knowledge_base

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/knowledge_base"
)

func (i *Implementation) DeleteKnowledgeBaseById(ctx context.Context, request *desc.DeleteKnowledgeBaseByIdRequest) (*desc.DeleteKnowledgeBaseByIdResponse, error) {
	knowledgeBaseID, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "knowledge base id is invalid")
	}

	err = i.knowledgeBaseService.DeleteById(ctx, knowledgeBaseID)
	if err != nil {
		return nil, status.Error(codes.Internal, "can't delete knowledge base")
	}

	return &desc.DeleteKnowledgeBaseByIdResponse{}, nil
}
