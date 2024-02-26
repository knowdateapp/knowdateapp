package collection

import (
	"google.golang.org/grpc"

	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/collection"
)

type Implementation struct {
	desc.UnimplementedCollectionServiceServer
}

func NewCollectionService() *Implementation {
	return &Implementation{}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.CollectionServiceServer) {
	desc.RegisterCollectionServiceServer(s, srv)
}
