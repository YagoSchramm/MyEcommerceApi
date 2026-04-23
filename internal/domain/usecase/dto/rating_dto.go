package dto

import (
	"time"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service"
	"github.com/google/uuid"
)

type CreateRatingDTO struct {
	UserID      uuid.UUID
	UserName    string    `json:"user_name" db:"user_name"`
	ProdutctID  uuid.UUID `json:"product_id" db:"product_id"`
	PurchaseID  uuid.UUID `json:"purchase_id" db:"purchase_id"`
	Description string    `json:"description" db:"description"`
	Rating      float32   `json:"rating" db:"rating"`
}
type UpdateRatingDTO struct {
	ID        uuid.UUID `json:"id" db:"rating_id"`
	Rating    float32   `json:"rating" db:"rating"`
	UpdatedAt time.Time `json:"updated_at"`
}
type DeleteRatingDTO struct {
	ID         uuid.UUID `json:"id" db:"rating_id"`
	UserID     uuid.UUID
	ProdutctID uuid.UUID  `json:"product_id" db:"product_id"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
type GetRatingByIdDTO struct {
	ID uuid.UUID `json:"id" db:"rating_id"`
}
type GetRatingByUserIdDTO struct {
	UserID uuid.UUID `json:"user_id" db:"user_id"`
}
type GetAllRatingByProductIdDTO struct {
	ProductID uuid.UUID `json:"product_id" db:"product_id"`
}
type RatingResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	UserName    string    `json:"user_name" `
	ProdutctID  uuid.UUID `json:"product_id" `
	PurchaseID  uuid.UUID `json:"purchase_id" `
	Rating      float32   `json:"rating" `
	Description string    `json:"description" `
}

func ToRatingEntity(dto CreateRatingDTO) *entity.Rating {
	return &entity.Rating{
		ID:          *service.GenerateID(),
		UserID:      dto.UserID,
		PurchaseID:  dto.PurchaseID,
		UserName:    dto.UserName,
		ProductID:   dto.ProdutctID,
		Description: dto.Description,
		DeletedAt:   nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
func ToRatingResponse(rating *entity.Rating) RatingResponseDTO {
	return RatingResponseDTO{
		ID:          rating.ID,
		UserName:    rating.UserName,
		ProdutctID:  rating.ProductID,
		PurchaseID:  rating.PurchaseID,
		Rating:      rating.Rating,
		Description: rating.Description,
	}
}
