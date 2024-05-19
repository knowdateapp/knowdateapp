package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/knowdateapp/knowdateapp/backend/user/internal/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
}

type UserService struct {
	repository UserRepository
	logger     *slog.Logger
}

func NewUserService(repository UserRepository, logger *slog.Logger) *UserService {
	return &UserService{
		repository: repository,
		logger:     logger,
	}
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	err := s.repository.Create(ctx, user)
	if err != nil {
		s.logger.Error(fmt.Sprintf("User %s not created: %s", user.Username, err))
		return err
	}
	return nil
}
