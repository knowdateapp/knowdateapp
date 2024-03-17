package entities

import (
	"time"

	"github.com/google/uuid"
)

type Collection struct {
	ID              uuid.UUID   `json:"id" bson:"id"`
	KnowledgeBaseID uuid.UUID   `json:"knowledge_base_id" bson:"knowledge_base_id"`
	OwnerID         uuid.UUID   `json:"owner_id" bson:"owner_id"`
	Topic           string      `json:"topic" bson:"topic"`
	Description     string      `json:"description" bson:"description"`
	TopicIds        []uuid.UUID `json:"topic_ids" bson:"topic_ids"`
	CreatedAt       time.Time   `json:"created_at" bson:"created_at"`
}
