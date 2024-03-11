package auth

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	desc "github.com/knowdateapp/knowdateapp/backend/user/internal/pb/api/v1/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var signingKey = []byte("test")

func (i *Implementation) LoginAnonymousUser(_ context.Context, request *desc.LoginAnonymousUserRequest) (*desc.LoginAnonymousUserResponse, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return nil, status.Error(codes.Internal, "token generation error")
	}
	return &desc.LoginAnonymousUserResponse{
		Token:        tokenString,
		RefreshToken: tokenString,
	}, nil
}
