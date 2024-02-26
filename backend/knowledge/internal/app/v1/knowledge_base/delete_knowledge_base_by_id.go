package knowledge_base

import (
	"context"

	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/knowledge_base"
)

func (i *Implementation) DeleteKnowledgeBaseById(_ context.Context, _ *desc.DeleteKnowledgeBaseByIdRequest) (*desc.DeleteKnowledgeBaseByIdResponse, error) {
	return &desc.DeleteKnowledgeBaseByIdResponse{}, nil
}
