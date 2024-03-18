package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/knowdateapp/knowdateapp/backend/user/internal/app/v1/auth"
	"github.com/knowdateapp/knowdateapp/backend/user/internal/app/v1/user"
)

func main() {
	const (
		ServerAddressGRPC = ":1028"
	)

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Can't create connection pool:", err)
	}
	defer pool.Close()

	var version string
	err = pool.QueryRow(context.Background(), "select version()").Scan(&version)
	if err != nil {
		log.Println("QueryRow failed:", err)
	}
	log.Println("version: ", version)

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

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		err = server.Serve(lis)
		if err != nil {
			log.Println("Server error:", err)
			close(stop)
		}
	}()

	<-stop
	server.GracefulStop()
	log.Println("Server stopped")
}
