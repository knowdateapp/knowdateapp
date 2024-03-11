package knowledge_base

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/common"
	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/knowledge_base"
)

func (i *Implementation) GetKnowledgeBaseById(_ context.Context, request *desc.GetKnowledgeBaseByIdRequest) (*desc.GetKnowledgeBaseByIdResponse, error) {
	return &desc.GetKnowledgeBaseByIdResponse{
		KnowledgeBase: &common.KnowledgeBase{
			Id:          request.Id,
			OwnerId:     uuid.New().String(),
			Topic:       "Topic",
			Description: "Description",
			Collections: nil,
			CreatedAt:   timestamppb.Now(),
		},
	}, nil
}
