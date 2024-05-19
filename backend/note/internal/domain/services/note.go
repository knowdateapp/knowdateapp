package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

type NoteRepository interface {
	Create(ctx context.Context, note *models.Note) (*models.Note, error)
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

func (s *NoteService) Create(ctx context.Context, note *models.Note) (*models.Note, error) {
	result, err := s.repository.Create(ctx, note)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Note not created: ownerId %s, title '%s', err %s", note.OwnerId, note.Title, err))
		return nil, err
	}
	return result, nil
}
