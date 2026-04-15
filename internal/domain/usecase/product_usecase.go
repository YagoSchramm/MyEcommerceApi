package usecase

import (
	"context"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/rules"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
	"github.com/google/uuid"
)

type ProductUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}
func (usc *ProductUsecase) CreateProduct(ctx context.Context, product *dto.CreateProductDTO) (*uuid.UUID, error) {
	err := rules.ValidateCreateProduct(*product)
	if err != nil {
		return nil, err
	}
	err = rules.ValidateImageFile(product.Image)
	if err != nil {
		return nil, err
	}
	productEntity := dto.ToProductEntity(*product)
	return usc.repo.CreateProduct(ctx, *productEntity)
}
func (usc *ProductUsecase) UpdateProduct(ctx context.Context, updateIt *dto.UpdateProductDTO) error {
	err := rules.ValidateUpdateProduct(*updateIt)
	if err != nil {
		return err
	}
	return usc.repo.UpdateProduct(ctx, *updateIt)
}
func (usc *ProductUsecase) DeleteProduct(ctx context.Context, deleteIt *dto.DeleteProductDTO) error {
	return usc.repo.DeleteProduct(ctx, *deleteIt)
}
func (usc *ProductUsecase) GetProductById(ctx context.Context, input *dto.GetProductByIdDTO) (*dto.ProductResponse, error) {
	productEntity, err := usc.repo.GetProductById(ctx, input.ID.String())
	if err != nil {
		return nil, err
	}
	product := dto.ToProductResponse(productEntity)
	return &product, nil
}
func (usc *ProductUsecase) GetAllProducts(ctx context.Context, input *dto.GetAllProductsDTO) ([]*dto.ProductResponse, error) {
	productEntities, err := usc.repo.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	var products []*dto.ProductResponse
	for _, productEntity := range productEntities {
		product := dto.ToProductResponse(productEntity)
		products = append(products, &product)
	}
	return products, nil
}
