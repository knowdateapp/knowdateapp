package collection

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/collection"
	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/common"
)

func (i *Implementation) CreateCollection(_ context.Context, request *desc.CreateCollectionRequest) (*desc.CreateCollectionResponse, error) {
	return &desc.CreateCollectionResponse{
		Collection: &common.Collection{
			Id:              uuid.New().String(),
			KnowledgeBaseId: request.KnowledgeBaseId,
			OwnerId:         request.OwnerId,
			Topic:           request.Topic,
			Description:     request.Description,
			TopicIds:        nil,
			CreatedAt:       timestamppb.Now(),
		},
	}, nil
}
