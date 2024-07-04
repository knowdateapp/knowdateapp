package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

type CardRepository interface {
	Create(ctx context.Context, card *models.Card) (string, error)
	Update(ctx context.Context, card *models.Card) error
	GetByID(ctx context.Context, workspace string, ID string) (*models.Card, error)
	GetByWorkspace(ctx context.Context, workspace string) ([]*models.Card, error)
}

type CardService struct {
	repository CardRepository
	logger     *slog.Logger
}

func NewCardService(repository CardRepository, logger *slog.Logger) *CardService {
	return &CardService{
		repository: repository,
		logger:     logger,
	}
}

func (s *CardService) Create(ctx context.Context, card *models.Card) (*models.Card, error) {
	id, err := s.repository.Create(ctx, card)
	if err != nil {
		s.logger.Error(fmt.Sprintf("card was not created: %s", err))
		return nil, err
	}

	result := &models.Card{
		ID:        id,
		Question:  card.Question,
		Answer:    card.Answer,
		Workspace: card.Workspace,
	}
	return result, nil
}

func (s *CardService) Update(ctx context.Context, card *models.Card) (*models.Card, error) {
	c, err := s.repository.GetByID(ctx, card.Workspace, card.ID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("card was not updated: %s", err))
		return nil, err
	}

	c.Question = card.Question
	c.Answer = card.Answer

	err = s.repository.Update(ctx, c)
	if err != nil {
		s.logger.Error(fmt.Sprintf("card was not updated: %s", err))
		return nil, err
	}

	return c, nil
}

func (s *CardService) GetByWorkspace(ctx context.Context, workspace string) ([]*models.Card, error) {
	cards, err := s.repository.GetByWorkspace(ctx, workspace)
	if err != nil {
		s.logger.Error(fmt.Sprintf("can not get cards by workspace: %s", err))
		return nil, err
	}
	return cards, nil
}

func (s *CardService) GetByID(ctx context.Context, workspace string, ID string) (*models.Card, error) {
	card, err := s.repository.GetByID(ctx, workspace, ID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("can not get the card by id: %s", err))
		return nil, err
	}
	return card, nil
}
