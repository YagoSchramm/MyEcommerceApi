package dto

import (
	"time"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/util"
	"github.com/google/uuid"
)

type CreatePurchaseDTO struct {
	ID        uuid.UUID `json:"purchase_id"`
	ProductID uuid.UUID `json:"product_id"`
	UserID    uuid.UUID
	Quantity  int `json:"quantity"`
}
type PurchaseResponseDTO struct {
	ID        uuid.UUID `json:"purchase_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

func ToPurchaseEntity(dto CreatePurchaseDTO) *entity.Purchase {
	return &entity.Purchase{
		ID:        *util.GenerateID(),
		ProductID: dto.ProductID,
		UserID:    dto.UserID,
		Quantity:  dto.Quantity,
		CreatedAt: time.Now(),
	}
}
func ToPurchaseRespone(purchase *entity.Purchase) PurchaseResponseDTO {
	return PurchaseResponseDTO{
		ID:        purchase.ID,
		ProductID: purchase.ProductID,
		Quantity:  purchase.Quantity,
	}
}
