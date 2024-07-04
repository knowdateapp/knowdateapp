package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/knowdateapp/knowdateapp/backend/note/internal/domain/models"
)

type CardRepository struct {
	db         *mongo.Client
	database   string
	collection string
}

func NewCardRepository(db *mongo.Client, database string, collection string) *CardRepository {
	return &CardRepository{
		db:         db,
		database:   database,
		collection: collection,
	}
}

func (r *CardRepository) Create(ctx context.Context, card *models.Card) (string, error) {
	result, err := r.db.Database(r.database).
		Collection(r.collection).
		InsertOne(ctx, card)
	if err != nil {
		return "", fmt.Errorf("failed to insert card document: %s", err)
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to insert card document: interted id has invalid format %v", result.InsertedID)
	}

	return id.Hex(), nil
}

func (r *CardRepository) Update(ctx context.Context, card *models.Card) error {
	objectID, err := primitive.ObjectIDFromHex(card.ID)
	if err != nil {
		return fmt.Errorf("failed to parse object id %s: %s", card.ID, err)
	}

	update := bson.D{{"$set", bson.D{{"question", card.Question}, {"answer", card.Answer}}}}

	result, err := r.db.Database(r.database).
		Collection(r.collection).
		UpdateByID(ctx, objectID, update)
	if err != nil {
		return fmt.Errorf("failed to update card: %s", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("failed to update card: not found")
	}

	return nil
}

func (r *CardRepository) GetByID(ctx context.Context, workspace string, ID string) (*models.Card, error) {
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse object id %s: %s", ID, err)
	}

	result := r.db.Database(r.database).
		Collection(r.collection).
		FindOne(ctx, bson.D{{"_id", objectID}, {"workspace", workspace}})

	card := &models.Card{}
	err = result.Decode(&card)
	if err != nil {
		return nil, fmt.Errorf("failed to find card by id: %s", err)
	}

	return card, nil
}

func (r *CardRepository) GetByWorkspace(ctx context.Context, workspace string) ([]*models.Card, error) {
	result, err := r.db.Database(r.database).
		Collection(r.collection).
		Find(ctx, bson.D{{"workspace", workspace}})
	if err != nil {
		return nil, fmt.Errorf("failed to find cards: %s", err)
	}

	var cards []*models.Card
	err = result.All(ctx, &cards)
	if err != nil {
		return nil, fmt.Errorf("failed to find cards: unmarshal error %s", err)
	}

	return cards, nil
}
