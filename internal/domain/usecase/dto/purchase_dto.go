package dto

import (
	"time"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service"
	"github.com/google/uuid"
)

type CreatePurchaseDTO struct {
	ProductID uuid.UUID `json:"product_id"`
	UserID    uuid.UUID
	Quantity  int `json:"quantity"`
}
type GetPurchaseByIdDTO struct {
	ID uuid.UUID `json:"purchase_id"`
}
type GetAllPurchaseByUserIdDTO struct {
	UserID uuid.UUID `json:"user_id"`
}
type GetAllPurchasesDTO struct {
	ID uuid.UUID `json:"purchase_id"`
}
type PurchaseResponseDTO struct {
	ID        uuid.UUID `json:"purchase_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

func ToPurchaseEntity(dto CreatePurchaseDTO, price float32) *entity.Purchase {
	return &entity.Purchase{
		ID:        *service.GenerateID(),
		ProductID: dto.ProductID,
		UserID:    dto.UserID,
		Value:     price,
		Quantity:  dto.Quantity,
		CreatedAt: time.Now(),
	}
}
func ToPurchaseResponse(purchase *entity.Purchase) PurchaseResponseDTO {
	return PurchaseResponseDTO{
		ID:        purchase.ID,
		ProductID: purchase.ProductID,
		Quantity:  purchase.Quantity,
	}
}
