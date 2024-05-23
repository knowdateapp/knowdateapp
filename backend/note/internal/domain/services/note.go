package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

type NoteRepository interface {
	Create(ctx context.Context, note *models.Note) (string, error)
	GetByWorkspace(ctx context.Context, workspace string) ([]*models.Note, error)
}

type NoteService struct {
	repository NoteRepository
	logger     *slog.Logger
}

func NewNoteService(repository NoteRepository, logger *slog.Logger) *NoteService {
	return &NoteService{
		repository: repository,
		logger:     logger,
	}
}

func (s *NoteService) Create(ctx context.Context, note *models.Note) (string, error) {
	id, err := s.repository.Create(ctx, note)
	if err != nil {
		s.logger.Error(fmt.Sprintf("note was not created: %s", err))
		return "", err
	}
	return id, nil
}

func (s *NoteService) GetByWorkspace(ctx context.Context, workspace string) ([]*models.Note, error) {
	notes, err := s.repository.GetByWorkspace(ctx, workspace)
	if err != nil {
		s.logger.Error(fmt.Sprintf("can not get notes by workspace: %s", err))
		return nil, err
	}
	return notes, nil
}
