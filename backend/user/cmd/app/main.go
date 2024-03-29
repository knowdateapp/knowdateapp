package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/knowdateapp/knowdateapp/backend/user/internal/app/v1/auth"
	"github.com/knowdateapp/knowdateapp/backend/user/internal/app/v1/user"
)

func main() {
	const (
		ServerAddressGRPC = ":1028"
	)

	// GRPC server setup.
	server := grpc.NewServer()

	auth.RegisterGRPCServer(server, auth.NewAuthService())
	user.RegisterGRPCServer(server, user.NewUserService())

	reflection.Register(server)

	// Start application.
	log.Println("User server startup on port:", ServerAddressGRPC)

	lis, err := net.Listen("tcp", ServerAddressGRPC)
	if err != nil {
		log.Fatalln("Can't start application:", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		err = server.Serve(lis)
		if err != nil {
			log.Println("Server error:", err)
		}
		stop()
	}()

	<-ctx.Done()
	server.GracefulStop()
	log.Println("Server stopped")
}
