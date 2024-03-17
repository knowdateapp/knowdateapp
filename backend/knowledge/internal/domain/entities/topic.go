package entities

import (
	"time"

	"github.com/google/uuid"
)

type Topic struct {
	ID              uuid.UUID `json:"id" bson:"id"`
	KnowledgeBaseID uuid.UUID `json:"knowledge_base_id" bson:"knowledge_base_id"`
	CollectionID    uuid.UUID `json:"collection_id" bson:"collection_id"`
	OwnerID         uuid.UUID `json:"owner_id" bson:"owner_id"`
	Topic           string    `json:"topic" bson:"topic"`
	Description     string    `json:"description" bson:"description"`
	Order           int32     `json:"order" bson:"order"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
}
