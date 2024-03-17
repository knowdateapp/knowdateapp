package knowledge_base

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/common"
	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/knowledge_base"
)

func (i *Implementation) CreateKnowledgeBase(ctx context.Context, request *desc.CreateKnowledgeBaseRequest) (*desc.CreateKnowledgeBaseResponse, error) {
	ownerID, err := uuid.Parse(request.OwnerId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "owner id is invalid")
	}

	knowledgeBase, err := i.knowledgeBaseService.Create(ctx, ownerID, request.Topic, request.Description)
	if err != nil {
		return nil, status.Error(codes.Internal, "knowledge base not created")
	}

	collections := make([]*common.Collection, 0, len(knowledgeBase.Collections))

	for _, collection := range knowledgeBase.Collections {
		topicIds := make([]string, 0, len(collection.TopicIds))

		for _, id := range collection.TopicIds {
			topicIds = append(topicIds, id.String())
		}

		collections = append(collections, &common.Collection{
			Id:              collection.ID.String(),
			KnowledgeBaseId: collection.KnowledgeBaseID.String(),
			OwnerId:         collection.OwnerID.String(),
			Topic:           collection.Topic,
			Description:     collection.Description,
			TopicIds:        topicIds,
			CreatedAt:       timestamppb.New(collection.CreatedAt),
		})
	}

	return &desc.CreateKnowledgeBaseResponse{
		KnowledgeBase: &common.KnowledgeBase{
			Id:          knowledgeBase.ID.String(),
			OwnerId:     knowledgeBase.OwnerID.String(),
			Topic:       knowledgeBase.Topic,
			Description: knowledgeBase.Description,
			Collections: collections,
			CreatedAt:   timestamppb.New(knowledgeBase.CreatedAt),
		},
	}, nil
}
