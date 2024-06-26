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
	objectID, err := primitive.ObjectIDFromHex(note.ID)
	if err != nil {
		return fmt.Errorf("failed to parse object id %s: %s", note.ID, err)
	}

	update := bson.D{{"$set", bson.D{{"title", note.Title}, {"content_uri", note.ContentUri}}}}

	result, err := r.db.Database(r.database).
		Collection(r.collection).
		UpdateByID(ctx, objectID, update)
	if err != nil {
		return fmt.Errorf("failed to update note: %s", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("failed to update note: not found")
	}

	return nil
}

func (r *NoteRepository) GetByID(ctx context.Context, workspace string, ID string) (*models.Note, error) {
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse object id %s: %s", ID, err)
	}

	result := r.db.Database(r.database).
		Collection(r.collection).
		FindOne(ctx, bson.D{{"_id", objectID}, {"workspace", workspace}})

	note := &models.Note{}
	err = result.Decode(&note)
	if err != nil {
		return nil, fmt.Errorf("failed to find not by id: %s", err)
	}

	return note, nil
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
