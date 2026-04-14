package service

import (
	"context"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/rules"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service/dto"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}
func (srv *ProductService) CreateProduct(ctx context.Context, product *dto.CreateProductDTO) error {
	err := rules.ValidateCreateProduct(*product)
	if err != nil {
		return err
	}
	productEntity := dto.ToProductEntity(*product)
	return srv.repo.CreateProduct(ctx, *productEntity)
}
func (srv *ProductService) UpdateProduct(ctx context.Context, updateIt *dto.UpdateProductDTO) error {
	err := rules.ValidateUpdateProduct(*updateIt)
	if err != nil {
		return err
	}
	return srv.repo.UpdateProduct(ctx, *updateIt)
}
func (srv *ProductService) DeleteProduct(ctx context.Context, deleteIt *dto.DeleteProductDTO) error {
	return srv.repo.DeleteProduct(ctx, *deleteIt)
}
func (srv *ProductService) GetProductById(ctx context.Context, input *dto.GetProductByIdDTO) (*dto.ProductResponse, error) {
	productEntity, err := srv.repo.GetProductById(ctx, input.ID.String())
	if err != nil {
		return nil, err
	}
	product := dto.ToProductResponse(productEntity)
	return &product, nil
}
func (srv *ProductService) GetAllProducts(ctx context.Context, input *dto.GetAllProductsDTO) ([]*dto.ProductResponse, error) {
	productEntities, err := srv.repo.GetAllProducts(ctx)
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
