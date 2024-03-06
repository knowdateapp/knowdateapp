package auth

import (
	"context"

	desc "github.com/knowdateapp/knowdateapp/backend/user/internal/pb/api/v1/auth"
)

func (i *Implementation) LoginAnonymousUser(_ context.Context, request *desc.LoginAnonymousUserRequest) (*desc.LoginAnonymousUserResponse, error) {
	return &desc.LoginAnonymousUserResponse{
		Token:        "qwerty",
		RefreshToken: "asdfg",
	}, nil
}
