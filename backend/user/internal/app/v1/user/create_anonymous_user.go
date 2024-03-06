package user

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/knowdateapp/knowdateapp/backend/user/internal/pb/api/v1/user"
)

func (i *Implementation) CreateAnonymousUser(_ context.Context, request *desc.CreateAnonymousUserRequest) (*desc.CreateAnonymousUserResponse, error) {
	return &desc.CreateAnonymousUserResponse{
		User: &desc.User{
			Id:        uuid.New().String(),
			Username:  request.Username,
			CreatedAt: timestamppb.Now(),
		},
	}, nil
}
