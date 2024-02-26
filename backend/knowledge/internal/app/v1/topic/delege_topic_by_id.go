package topic

import (
	"context"

	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/topic"
)

func (i *Implementation) DeleteTopicById(_ context.Context, _ *desc.DeleteTopicByIdRequest) (*desc.DeleteTopicByIdResponse, error) {
	return &desc.DeleteTopicByIdResponse{}, nil
}
