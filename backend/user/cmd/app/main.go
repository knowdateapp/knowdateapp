package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/knowdateapp/knowdateapp/backend/user/internal/app/v1/auth"
	"github.com/knowdateapp/knowdateapp/backend/user/internal/app/v1/user"
)

func main() {
	const (
		ServerAddressGRPC = ":90"
	)

	// GRPC server setup.
	server := grpc.NewServer()

	auth.RegisterGRPCServer(server, auth.NewAuthService())
	user.RegisterGRPCServer(server, user.NewUserService())

	reflection.Register(server)

	// Start application.
	log.Println("User server startup on port:", ServerAddressGRPC)

	go func() {
		lis, err := net.Listen("tcp", ServerAddressGRPC)
		if err != nil {
			log.Fatalln("Can't start application:", err)
		}

		err = server.Serve(lis)
		if err != nil {
			log.Fatalln("Can't start application:", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	server.GracefulStop()
	log.Println("Server stopped")
}
