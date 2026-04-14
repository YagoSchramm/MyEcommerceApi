package service

import (
	"context"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/rules"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service/dto"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
)

type PurchaseService struct {
	repo *repository.PurchaseRepository
}

func NewPurchaseService(repo *repository.PurchaseRepository) *PurchaseService {
	return &PurchaseService{repo: repo}
}
func (srv *PurchaseService) CreatePurchase(ctx context.Context, purchase *dto.CreatePurchaseDTO) error {
	price, err := srv.repo.GetPriceByProductId(ctx, purchase.ProductID.String())
	if err != nil {
		return err
	}
	value := price * float32(purchase.Quantity)
	err = rules.ValidateCreatePurchase(*purchase)
	if err != nil {
		return err
	}
	purchaseEntity := dto.ToPurchaseEntity(*purchase, value)
	return srv.repo.CreatePurchase(ctx, *purchaseEntity)
}
