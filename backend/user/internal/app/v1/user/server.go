package user

import (
	"google.golang.org/grpc"

	desc "github.com/knowdateapp/knowdateapp/backend/user/internal/pb/api/v1/user"
)

type Implementation struct {
	desc.UnimplementedUserServiceServer
}

func NewUserService() *Implementation {
	return &Implementation{}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.UserServiceServer) {
	desc.RegisterUserServiceServer(s, srv)
}
