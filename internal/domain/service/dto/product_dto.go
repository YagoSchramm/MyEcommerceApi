package dto

import (
	"time"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/util"
	"github.com/google/uuid"
)

type CreateProductDTO struct {
	UserID      uuid.UUID `json:"user_id" db:"user_id" validate:"required"`
	UserName    string    `json:"user_name" db:"user_name" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Value       float32   `json:"value" validate:"required"`
	Image       string    `json:"image" validate:"required"`
	Stock       int32     `json:"stock" validate:"required"`
	Description string    `json:"description" validate:"required"`
}
type UpdateProductDTO struct {
	ID          uuid.UUID
	Name        string  `json:"name" validate:"required"`
	Value       float32 `json:"value" validate:"required"`
	Image       string  `json:"image" validate:"required"`
	Stock       int32   `json:"stock" validate:"required"`
	Description string  `json:"description" validate:"required"`
}
type DeleteProductDTO struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ProductID uuid.UUID
}
type GetProductByIdDTO struct {
	ID uuid.UUID
}
type GetAllProductsDTO struct {
	ID uuid.UUID
}

type ProductResponse struct {
	ID           uuid.UUID `json:"product_id"`
	UserName     string    `json:"user_name"`
	Name         string    `json:"name"`
	Value        float32   `json:"value"`
	Image        string    `json:"image"`
	Stock        int32     `json:"stock"`
	Description  string    `json:"description"`
	AvgRating    float32   `json:"avg_rating"`
	TotalReviews int32     `json:"total_reviews"`
}

func ToProductEntity(dto CreateProductDTO) *entity.Product {
	return &entity.Product{
		ID:          *util.GenerateID(),
		UserID:      dto.UserID,
		UserName:    dto.UserName,
		Name:        dto.Name,
		Image:       dto.Image,
		Stock:       dto.Stock,
		Value:       dto.Value,
		Description: dto.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
}
func ToProductResponse(product *entity.Product) ProductResponse {
	return ProductResponse{
		ID:           product.ID,
		UserName:     product.UserName,
		Name:         product.Name,
		Image:        product.Image,
		Value:        product.Value,
		Stock:        product.Stock,
		Description:  product.Description,
		AvgRating:    *product.AvgRating,
		TotalReviews: product.TotalRatings,
	}
}
