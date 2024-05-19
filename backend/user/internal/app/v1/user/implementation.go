package user

import (
	"context"

	desc "github.com/knowdateapp/knowdateapp/backend/user/internal/api/http/v1/user"
	"github.com/knowdateapp/knowdateapp/backend/user/internal/domain/models"
)

type Service interface {
	Create(ctx context.Context, user *models.User) error
}

type Implementation struct {
	desc.Unimplemented

	service Service
}

func NewUserServerImplementation(service Service) *Implementation {
	return &Implementation{
		service: service,
	}
}
