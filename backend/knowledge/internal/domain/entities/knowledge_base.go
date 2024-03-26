package entities

import (
	"time"

	"github.com/google/uuid"
)

type KnowledgeBase struct {
	ID          uuid.UUID    `json:"id" bson:"id"`
	OwnerID     uuid.UUID    `json:"owner_id" bson:"owner_id"`
	Topic       string       `json:"topic" bson:"topic"`
	Description string       `json:"description" bson:"description"`
	Collections []Collection `json:"collections" bson:"collections"`
	CreatedAt   time.Time    `json:"created_at" bson:"created_at"`
}
