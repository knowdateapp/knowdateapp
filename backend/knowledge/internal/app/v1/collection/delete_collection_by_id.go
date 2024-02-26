package collection

import (
	"context"

	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/collection"
)

func (i *Implementation) DeleteCollectionById(_ context.Context, _ *desc.DeleteCollectionByIdRequest) (*desc.DeleteCollectionByIdResponse, error) {
	return &desc.DeleteCollectionByIdResponse{}, nil
}
