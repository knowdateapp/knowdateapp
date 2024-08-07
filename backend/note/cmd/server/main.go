package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	carddesc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/card"
	notedesc "github.com/knowdateapp/knowdateapp/backend/note/internal/api/http/v1/note"
	cardapp "github.com/knowdateapp/knowdateapp/backend/note/internal/app/v1/card"
	noteapp "github.com/knowdateapp/knowdateapp/backend/note/internal/app/v1/note"
	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/services"
	"github.com/knowdateapp/knowdateapp/backend/note/internal/infrastructure/repositories"
)

func main() {
	const (
		ServerAddress      = ":80"
		DebugServerAddress = ":84"
		BaseURL            = "/api/note/v1"
	)

	// App configuration.
	var (
		dbUsername  = flag.String("db-username", "admin", "username to connect to the database")
		dbPassword  = flag.String("db-password", "password", "password to connect to the database")
		dbAddress   = flag.String("db-address", "localhost:27017", "database host and port")
		storagePath = flag.String("storage-path", "storage", "service storage path")
	)

	flag.Parse()

	ctx := context.Background()

	// Logger configuration.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))

	// Configure MongoDB connection.
	mongoConnectionURI := fmt.Sprintf("mongodb://%s:%s@%s/", *dbUsername, *dbPassword, *dbAddress)
	mongoClientOptions := options.Client().
		ApplyURI(mongoConnectionURI).
		SetMinPoolSize(10).
		SetMaxPoolSize(100).
		SetMaxConnIdleTime(15 * time.Minute).
		SetTimeout(5 * time.Second)
	err := mongoClientOptions.Validate()
	if err != nil {
		logger.Error(fmt.Sprint("mongo client options are invalid:", err))
		os.Exit(1)
	}

	db, err := mongo.Connect(ctx, mongoClientOptions)
	if err != nil {
		logger.Error(fmt.Sprint("can not connect to MongoDB:", err))
		os.Exit(1)
	}

	// Cancel MongoDB connection.
	defer func() {
		dbDisconnectCtx, dbDisconnectCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer dbDisconnectCancel()

		if dErr := db.Disconnect(dbDisconnectCtx); dErr != nil {
			logger.Error(fmt.Sprint("the database client can not be disconnected:", dErr))
		}
	}()

	var (
		database        = "note"
		notesCollection = "notes"
		cardsCollection = "cards"
	)

	// Check MongoDB connection status.
	result := bson.M{}
	err = db.Database(database).RunCommand(ctx, bson.D{{"ping", 1}}).Decode(&result)
	if err != nil {
		logger.Error("the connection to the database has not been established:", err)
		os.Exit(1)
	}

	// Services setup.
	noteRepository := repositories.NewNoteRepository(db, database, notesCollection)
	noteStorage := repositories.NewNoteStorageFile(*storagePath)
	cardRepository := repositories.NewCardRepository(db, database, cardsCollection)

	noteService := services.NewNoteService(noteRepository, noteStorage, logger)
	cardService := services.NewCardService(cardRepository, logger)

	// HTTP server setup.
	noteServer := noteapp.NewNoteServerImplementation(noteService)
	cardServer := cardapp.NewCardServerImplementation(cardService)

	corsHandler := cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	// Start application at debug port.
	go func() {
		debugRouter := chi.NewRouter()
		debugRouter.Use(corsHandler)
		debugRouter.Handle(BaseURL+"/*", http.StripPrefix(BaseURL+"/", http.FileServer(http.Dir("./api/v1/"))))

		logger.Info(fmt.Sprint("start note debug service on port:", DebugServerAddress))

		if dErr := http.ListenAndServe(DebugServerAddress, debugRouter); dErr != nil {
			logger.Error(fmt.Sprint("can not run debug application:", dErr))
			os.Exit(1)
		}
	}()

	// Start application.
	router := chi.NewRouter()
	router.Use(corsHandler)

	notedesc.RegisterNoteServerHandler(router, noteServer, BaseURL)
	carddesc.RegisterCardServerHandler(router, cardServer, BaseURL)

	logger.Info(fmt.Sprint("note server startup on port:", ServerAddress))

	if err = http.ListenAndServe(ServerAddress, router); err != nil {
		logger.Error(fmt.Sprint("can not run application:", err))
		os.Exit(1)
	}

}
