package service

import (
	"context"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/rules"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service/dto"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
	"github.com/google/uuid"
)

type PurchaseService struct {
	repo *repository.PurchaseRepository
}

func NewPurchaseService(repo *repository.PurchaseRepository) *PurchaseService {
	return &PurchaseService{repo: repo}
}
func (srv *PurchaseService) CreatePurchase(ctx context.Context, purchase *dto.CreatePurchaseDTO) (*uuid.UUID, error) {
	price, err := srv.repo.GetPriceByProductId(ctx, purchase.ProductID.String())
	if err != nil {
		return nil, err
	}
	value := price * float32(purchase.Quantity)
	err = rules.ValidateCreatePurchase(*purchase)
	if err != nil {
		return nil, err
	}
	purchaseEntity := dto.ToPurchaseEntity(*purchase, value)
	return srv.repo.CreatePurchase(ctx, *purchaseEntity)
}
func (srv *PurchaseService) GetPurchaseById(ctx context.Context, input *dto.GetPurchaseByIdDTO) (*dto.PurchaseResponseDTO, error) {
	purchaseEntity, err := srv.repo.GetPurchaseById(ctx, input.ID.String())
	if err != nil {
		return nil, err
	}
	purchase := dto.ToPurchaseResponse(purchaseEntity)
	return &purchase, nil
}
func (srv *PurchaseService) GetAllPurchaseByUserId(ctx context.Context, input *dto.GetAllPurchaseByUserIdDTO) ([]*dto.PurchaseResponseDTO, error) {
	purchaseList, err := srv.repo.GetAllPurchaseByUserId(ctx, input.UserID.String())
	if err != nil {
		return nil, err
	}
	var purchaseListResponse []*dto.PurchaseResponseDTO
	for _, purchase := range purchaseList {
		purchaseResp := dto.ToPurchaseResponse(purchase)
		purchaseListResponse = append(purchaseListResponse, &purchaseResp)
	}
	return purchaseListResponse, err
}
func (srv *PurchaseService) GetAllPurchases(ctx context.Context, input *dto.GetAllPurchasesDTO) ([]*dto.PurchaseResponseDTO, error) {
	purchaseList, err := srv.repo.GetAllPurchases(ctx)
	if err != nil {
		return nil, err
	}
	var purchaseListResponse []*dto.PurchaseResponseDTO
	for _, purchase := range purchaseList {
		purchaseResp := dto.ToPurchaseResponse(purchase)
		purchaseListResponse = append(purchaseListResponse, &purchaseResp)
	}
	return purchaseListResponse, nil
}
