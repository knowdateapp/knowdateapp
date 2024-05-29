package services

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"

	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

type NoteRepository interface {
	Create(ctx context.Context, note *models.Note) (string, error)
	Update(ctx context.Context, note *models.Note) error
	Get(ctx context.Context, workspace string, ID string) (*models.Note, error)
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

func (s *NoteService) Create(ctx context.Context, note *models.Note) (*models.Note, error) {
	id, err := s.repository.Create(ctx, note)
	if err != nil {
		s.logger.Error(fmt.Sprintf("note was not created: %s", err))
		return nil, err
	}

	// TODO: Add default content uri.
	result := &models.Note{
		ID:         id,
		Title:      note.Title,
		Workspace:  note.Workspace,
		ContentUri: note.ContentUri,
	}
	return result, nil
}

func (s *NoteService) Update(ctx context.Context, note *models.Note, filename string, file *bytes.Buffer) (*models.Note, error) {
	// TODO: implement
	return nil, nil
}

func (s *NoteService) GetByWorkspace(ctx context.Context, workspace string) ([]*models.Note, error) {
	notes, err := s.repository.GetByWorkspace(ctx, workspace)
	if err != nil {
		s.logger.Error(fmt.Sprintf("can not get notes by workspace: %s", err))
		return nil, err
	}
	return notes, nil
}
