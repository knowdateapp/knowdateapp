package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/app/v1/collection"
	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/app/v1/knowledge_base"
	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/app/v1/topic"
	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/domain/services"
)

func main() {
	const (
		ServerAddressGRPC = ":82"
	)

	var (
		databaseUsername = flag.String("database-username", "admin", "username to connect to the database")
		databasePassword = flag.String("database-password", "password", "password to connect to the database")
		databaseAddress  = flag.String("database-address", "localhost", "database host and port")
	)

	flag.Parse()

	ctx := context.Background()

	// Configure MongoDB connection.
	mongoConnectionURI := fmt.Sprintf("mongodb://%s:%s@%s/", *databaseUsername, *databasePassword, *databaseAddress)
	mongoClientOptions := options.Client().
		ApplyURI(mongoConnectionURI).
		SetMinPoolSize(10).
		SetMaxPoolSize(100).
		SetMaxConnIdleTime(15 * time.Minute).
		SetTimeout(5 * time.Second)
	err := mongoClientOptions.Validate()
	if err != nil {
		log.Fatalln("MongoDB client options are invalid:", err)
	}

	db, err := mongo.Connect(ctx, mongoClientOptions)
	if err != nil {
		log.Fatalln("Can't connect to MongoDB:", err)
	}

	defer func() {
		dbDisconnectCtx, dbDisconnectCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer dbDisconnectCancel()

		if err = db.Disconnect(dbDisconnectCtx); err != nil {
			log.Println("The database can't be disconnected:", err)
		}
	}()

	result := bson.M{}
	err = db.Database("knowledge_db").
		RunCommand(ctx, bson.D{{"ping", 1}}).
		Decode(&result)
	if err != nil {
		log.Fatalln("The connection to the database has not been established:", err)
	}

	log.Printf("The connection to the database has been established: user=%s, address=%s", *databaseUsername, *databaseAddress)

	// Services setup.
	knowledgeBaseService := services.NewKnowledgeBaseService()

	// GRPC server setup.
	server := grpc.NewServer()

	knowledge_base.RegisterGRPCServer(server, knowledge_base.NewKnowledgeBaseService(knowledgeBaseService))
	collection.RegisterGRPCServer(server, collection.NewCollectionService())
	topic.RegisterGRPCServer(server, topic.NewTopicService())

	reflection.Register(server)

	// Start application.
	log.Println("Knowledge server startup on port:", ServerAddressGRPC)

	lis, err := net.Listen("tcp", ServerAddressGRPC)
	if err != nil {
		log.Fatalln("Can't start application:", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Fatalln("Can't start application:", err)
	}
}
