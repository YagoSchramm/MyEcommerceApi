package usecase

import (
	"context"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/rules"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
	"github.com/google/uuid"
)

type PurchaseUsecase struct {
	repo *repository.PurchaseRepository
}

func NewPurchaseUsecase(repo *repository.PurchaseRepository) *PurchaseUsecase {
	return &PurchaseUsecase{repo: repo}
}
func (usc *PurchaseUsecase) CreatePurchase(ctx context.Context, purchase *dto.CreatePurchaseDTO) (*uuid.UUID, error) {
	price, err := usc.repo.GetPriceByProductId(ctx, purchase.ProductID.String())
	if err != nil {
		return nil, err
	}
	value := price * float32(purchase.Quantity)
	err = rules.ValidateCreatePurchase(*purchase)
	if err != nil {
		return nil, err
	}
	purchaseEntity := dto.ToPurchaseEntity(*purchase, value)
	return usc.repo.CreatePurchase(ctx, *purchaseEntity)
}
func (usc *PurchaseUsecase) GetPurchaseById(ctx context.Context, input *dto.GetPurchaseByIdDTO) (*dto.PurchaseResponseDTO, error) {
	purchaseEntity, err := usc.repo.GetPurchaseById(ctx, input.ID.String())
	if err != nil {
		return nil, err
	}
	purchase := dto.ToPurchaseResponse(purchaseEntity)
	return &purchase, nil
}
func (usc *PurchaseUsecase) GetAllPurchaseByUserId(ctx context.Context, input *dto.GetAllPurchaseByUserIdDTO) ([]*dto.PurchaseResponseDTO, error) {
	purchaseList, err := usc.repo.GetAllPurchaseByUserId(ctx, input.UserID.String())
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
func (usc *PurchaseUsecase) GetAllPurchases(ctx context.Context, input *dto.GetAllPurchasesDTO) ([]*dto.PurchaseResponseDTO, error) {
	purchaseList, err := usc.repo.GetAllPurchases(ctx)
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
