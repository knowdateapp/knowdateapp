package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/knowdateapp/knowdateapp/backend/user/internal/domain/models"
)

type UserRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return nil
}
