package auth

import (
	"context"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/knowdateapp/knowdateapp/backend/user/internal/pb/api/v1/auth"
)

var signingKey = []byte("test")

func (i *Implementation) LoginAnonymousUser(_ context.Context, request *desc.LoginAnonymousUserRequest) (*desc.LoginAnonymousUserResponse, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": request.Id,
	})
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		log.Println("Auth token generation error:", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &desc.LoginAnonymousUserResponse{
		Token:        tokenString,
		RefreshToken: tokenString,
	}, nil
}
