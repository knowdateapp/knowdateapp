package services

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/domain/entities"
)

type knowledgeBaseRepositoryInterface interface {
	Create(ctx context.Context, knowledgeBase *entities.KnowledgeBase) error
	GetById(ctx context.Context, id uuid.UUID) (*entities.KnowledgeBase, error)
}

type KnowledgeBaseService struct {
	knowledgeBaseRepository knowledgeBaseRepositoryInterface
}

func NewKnowledgeBaseService(knowledgeBaseRepository knowledgeBaseRepositoryInterface) *KnowledgeBaseService {
	return &KnowledgeBaseService{
		knowledgeBaseRepository: knowledgeBaseRepository,
	}
}

func (s *KnowledgeBaseService) Create(
	ctx context.Context,
	ownerID uuid.UUID,
	topic string,
	description string,
) (*entities.KnowledgeBase, error) {
	knowledgeBase := &entities.KnowledgeBase{
		ID:          uuid.New(),
		OwnerID:     ownerID,
		Topic:       topic,
		Description: description,
		Collections: []entities.Collection{},
		CreatedAt:   time.Now(),
	}

	err := s.knowledgeBaseRepository.Create(ctx, knowledgeBase)
	if err != nil {
		return nil, err
	}

	return knowledgeBase, nil
}
