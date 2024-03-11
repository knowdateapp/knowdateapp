package topic

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/common"
	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/topic"
)

func (i *Implementation) CreateTopic(_ context.Context, request *desc.CreateTopicRequest) (*desc.CreateTopicResponse, error) {
	return &desc.CreateTopicResponse{
		Topic: &common.Topic{
			Id:              uuid.New().String(),
			KnowledgeBaseId: request.KnowledgeBaseId,
			CollectionId:    request.CollectionId,
			OwnerId:         request.OwnerId,
			Topic:           request.Topic,
			Description:     request.Description,
			Order:           1,
			CreatedAt:       timestamppb.Now(),
		},
	}, nil
}
