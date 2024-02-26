package topic

import (
	"google.golang.org/grpc"

	desc "github.com/knowdateapp/knowdateapp/backend/knowledge/internal/pb/api/v1/topic"
)

type Implementation struct {
	desc.UnimplementedTopicServiceServer
}

func NewTopicService() *Implementation {
	return &Implementation{}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.TopicServiceServer) {
	desc.RegisterTopicServiceServer(s, srv)
}
