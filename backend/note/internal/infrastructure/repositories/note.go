package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

type NoteRepository struct {
	db         *mongo.Client
	database   string
	collection string
}

func NewNoteRepository(db *mongo.Client, database string, collection string) *NoteRepository {
	return &NoteRepository{
		db:         db,
		database:   database,
		collection: collection,
	}
}

func (r *NoteRepository) Create(ctx context.Context, note *models.Note) (string, error) {
	result, err := r.db.Database(r.database).
		Collection(r.collection).
		InsertOne(ctx, note)
	if err != nil {
		return "", fmt.Errorf("failed to insert note document: %s", err)
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to insert note document: interted id has invalid format %v", result.InsertedID)
	}

	return id.Hex(), nil
}

func (r *NoteRepository) Update(ctx context.Context, note *models.Note) error {
	// TODO: implement
	return nil
}

func (r *NoteRepository) Get(ctx context.Context, workspace string, ID string) (*models.Note, error) {
	// TODO: implement
	return nil, nil
}

func (r *NoteRepository) GetByWorkspace(ctx context.Context, workspace string) ([]*models.Note, error) {
	result, err := r.db.Database(r.database).
		Collection(r.collection).
		Find(ctx, bson.D{{"workspace", workspace}})
	if err != nil {
		return nil, fmt.Errorf("failed to find notes: %s", err)
	}

	var notes []*models.Note
	err = result.All(ctx, &notes)
	if err != nil {
		return nil, fmt.Errorf("failed to find notes: unmarshal error %s", err)
	}

	return notes, nil
}
