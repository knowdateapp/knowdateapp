package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/knowdateapp/knowdateapp/backend/knowledge/internal/domain/entities"
)

type KnowledgeBaseRepository struct {
	// TODO: add request timeout duration
	dbName         string
	collectionName string
	db             *mongo.Client
}

func NewKnowledgeBaseRepository(dbName string, collectionName string, db *mongo.Client) *KnowledgeBaseRepository {
	return &KnowledgeBaseRepository{
		dbName:         dbName,
		collectionName: collectionName,
		db:             db,
	}
}

func (r *KnowledgeBaseRepository) Create(ctx context.Context, knowledgeBase *entities.KnowledgeBase) error {
	_, err := r.db.Database(r.dbName).Collection(r.collectionName).InsertOne(ctx, knowledgeBase)
	if err != nil {
		return fmt.Errorf("can't create a knowledge base: %w", err)
	}
	return nil
}

func (r *KnowledgeBaseRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.KnowledgeBase, error) {
	var (
		knowledgeBase = entities.KnowledgeBase{}
		filter        = bson.M{"knowledge_base_id": id.String()}
	)

	err := r.db.Database(r.dbName).
		Collection(r.collectionName).
		FindOne(ctx, filter).
		Decode(&knowledgeBase)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// TODO: return domain not found error
		}
		return nil, err
	}

	return &knowledgeBase, nil
}

func (r *KnowledgeBaseRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"knowledge_base_id": id.String()}
	_, err := r.db.Database(r.dbName).Collection(r.collectionName).DeleteOne(ctx, filter)
	return err
}
