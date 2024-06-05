package note

import (
	"bytes"
	"context"

	desc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/note"
	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

type Service interface {
	Create(ctx context.Context, note *models.Note) (*models.Note, error)
	Update(ctx context.Context, note *models.Note, filename string, file *bytes.Buffer) (*models.Note, error)
	GetByWorkspace(ctx context.Context, workspace string) ([]*models.Note, error)
}

type Implementation struct {
	desc.Unimplemented

	service Service
}

func NewNoteServerImplementation(service Service) *Implementation {
	return &Implementation{
		service: service,
	}
}
