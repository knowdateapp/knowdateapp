package card

import (
	"context"

	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/card"
	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

type Service interface {
	Create(ctx context.Context, card *models.Card) (*models.Card, error)
	Update(ctx context.Context, card *models.Card) (*models.Card, error)
	GetByWorkspace(ctx context.Context, workspace string) ([]*models.Card, error)
	GetByID(ctx context.Context, workspace string, ID string) (*models.Card, error)
}

type Implementation struct {
	desc.Unimplemented

	service Service
}

func NewCardServerImplementation(service Service) *Implementation {
	return &Implementation{
		service: service,
	}
}
