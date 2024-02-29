package collection

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/collection"
	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/common"
)

func (i *Implementation) GetCollectionById(_ context.Context, request *desc.GetCollectionByIdRequest) (*desc.GetCollectionByIdResponse, error) {
	return &desc.GetCollectionByIdResponse{
		Collection: &common.Collection{
			Id:              request.CollectionId,
			KnowledgeBaseId: request.KnowledgeBaseId,
			OwnerId:         uuid.New().String(),
			Topic:           "Topic",
			Description:     "Description",
			TopicIds:        nil,
			CreatedAt:       timestamppb.Now(),
		},
	}, nil
}
