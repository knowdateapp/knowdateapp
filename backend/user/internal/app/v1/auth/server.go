package auth

import (
	"google.golang.org/grpc"

	desc "github.com/knowdateapp/knowdateapp/backend/user/internal/pb/api/v1/auth"
)

type Implementation struct {
	desc.UnimplementedAuthServiceServer
}

func NewAuthService() *Implementation {
	return &Implementation{}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.AuthServiceServer) {
	desc.RegisterAuthServiceServer(s, srv)
}
