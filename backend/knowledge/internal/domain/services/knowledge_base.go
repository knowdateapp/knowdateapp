package services

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/domain/entities"
)

type KnowledgeBaseService struct {
}

func NewKnowledgeBaseService() *KnowledgeBaseService {
	return &KnowledgeBaseService{}
}

func (s *KnowledgeBaseService) Create(
	ctx context.Context,
	ownerID uuid.UUID,
	topic string,
	description string,
) (*entities.KnowledgeBase, error) {
	entity := &entities.KnowledgeBase{
		ID:          uuid.New(),
		OwnerID:     ownerID,
		Topic:       topic,
		Description: description,
		Collections: []entities.Collection{},
		CreatedAt:   time.Now(),
	}

	return entity, nil
}
