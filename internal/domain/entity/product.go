package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID  `json:"product_id" db:"product_id"`
	UserID      uuid.UUID  `json:"user_id" db:"user_id" validate:"required"`
	UserName    uuid.UUID  `json:"user_name" db:"user_name" validate:"required"`
	Name        string     `json:"name" validate:"required"`
	Value       float32    `json:"value" validate:"required"`
	Image       string     `json:"image" validate:"required"`
	Stock       int32      `json:"stock" validate:"required"`
	Description string     `json:"description" validate:"required"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
