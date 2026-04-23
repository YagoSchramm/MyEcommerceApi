package entity

import (
	"time"

	"github.com/google/uuid"
)

type Rating struct {
	ID          uuid.UUID  `json:"id" db:"rating_id"`
	UserID      uuid.UUID  `json:"user_id" db:"user_id"`
	UserName    string     `json:"user_name" db:"user_name"`
	PurchaseID  uuid.UUID  `json:"purchase_id" db:"purchase_id"`
	ProductID   uuid.UUID  `json:"product_id" db:"product_id"`
	Rating      float32    `json:"rating" db:"rating"`
	Description string     `json:"description" db:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
