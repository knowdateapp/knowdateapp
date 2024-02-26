package knowledge_base

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/common"
	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/knowledge_base"
)

func (i *Implementation) CreateKnowledgeBase(_ context.Context, request *desc.CreateKnowledgeBaseRequest) (*desc.CreateKnowledgeBaseResponse, error) {
	return &desc.CreateKnowledgeBaseResponse{
		KnowledgeBase: &common.KnowledgeBase{
			Id:          uuid.New().String(),
			OwnerId:     request.OwnerId,
			Topic:       request.Topic,
			Description: request.Description,
			Collections: nil,
			CreatedAt:   timestamppb.Now(),
		},
	}, nil
}
