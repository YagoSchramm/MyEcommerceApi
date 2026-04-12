package entity

import (
	"time"

	"github.com/google/uuid"
)

type Purchase struct {
	ID        uuid.UUID `json:"purchase_id"`
	ProductID uuid.UUID `json:"product_id"`
	UserID    uuid.UUID `json:"user_id"`
	Value     float32   `json:"value"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}
