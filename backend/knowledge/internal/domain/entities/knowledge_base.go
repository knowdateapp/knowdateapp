package entities

import (
	"time"

	"github.com/google/uuid"
)

type KnowledgeBase struct {
	// FIXME: UUID stored in MongoDB in binary format. It is not convenient to use.
	ID          uuid.UUID    `json:"id" bson:"knowledge_base_id,omitempty"`
	OwnerID     uuid.UUID    `json:"owner_id" bson:"owner_id,omitempty"`
	Topic       string       `json:"topic" bson:"topic,omitempty"`
	Description string       `json:"description" bson:"description,omitempty"`
	Collections []Collection `json:"collections" bson:"collections,omitempty"`
	CreatedAt   time.Time    `json:"created_at" bson:"created_at,omitempty"`
}
