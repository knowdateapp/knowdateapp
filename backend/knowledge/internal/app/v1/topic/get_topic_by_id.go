package topic

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/common"
	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/topic"
)

func (i *Implementation) GetTopicById(_ context.Context, request *desc.GetTopicByIdRequest) (*desc.GetTopicByIdResponse, error) {
	return &desc.GetTopicByIdResponse{
		Topic: &common.Topic{
			Id:              uuid.New().String(),
			KnowledgeBaseId: request.KnowledgeBaseId,
			CollectionId:    request.CollectionId,
			OwnerId:         uuid.New().String(),
			Topic:           "Topic",
			Description:     "Description",
			Order:           1,
			CreatedAt:       timestamppb.Now(),
		},
	}, nil
}
