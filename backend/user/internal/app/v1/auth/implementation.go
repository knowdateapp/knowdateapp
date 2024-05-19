package auth

import (
	"context"

	desc "github.com/knowdateapp/knowdateapp/backend/user/internal/api/http/v1/auth"
)

type Service interface {
	CreateToken(ctx context.Context, username string, password string) (string, error)
}

type Implementation struct {
	desc.Unimplemented

	service Service
}

func NewAuthServerImplementation(service Service) *Implementation {
	return &Implementation{
		service: service,
	}
}
