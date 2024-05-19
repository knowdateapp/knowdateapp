package services

import (
	"context"
	"log/slog"
)

type AuthService struct {
	repository UserRepository
	logger     *slog.Logger
}

func NewAuthService(repository UserRepository, logger *slog.Logger) *AuthService {
	return &AuthService{
		repository: repository,
		logger:     logger,
	}
}

func (s AuthService) CreateToken(ctx context.Context, username string, password string) (string, error) {
	return "jwt-token-long-string", nil
}
