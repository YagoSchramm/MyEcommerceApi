package domain

import (
	"context"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/service/dto"
)

type DomainService interface {
	UserService
	ProductService
	PurchaseService
	RatingService
}
type UserService interface {
	CreateUser(ctx context.Context, user *dto.CreateUserDTO) error
	UpdateUser(ctx context.Context, updateIt *dto.UpdateUserDTO) error
	DeleteUser(ctx context.Context, deleteIt *dto.DeleteUserDTO) error
	GetUserById(ctx context.Context, input *dto.GetUserByIdDTO) (*dto.UserResponseDTO, error)
	GetUserByRole(ctx context.Context, input *dto.GetUserByRoleDTO) ([]*dto.UserResponseDTO, error)
	GetAllUsers(ctx context.Context, input *dto.GetAllUsersDTO) ([]*dto.UserResponseDTO, error)
}
type ProductService interface {
	CreateProduct(ctx context.Context, product *dto.CreateProductDTO) error
	UpdateProduct(ctx context.Context, updateIt *dto.UpdateProductDTO) error
	DeleteProduct(ctx context.Context, deleteIt *dto.DeleteProductDTO) error
	GetProductById(ctx context.Context, input *dto.GetProductByIdDTO) (*dto.ProductResponse, error)
	GetAllProducts(ctx context.Context, input *dto.GetAllProductsDTO) ([]*dto.ProductResponse, error)
}
type PurchaseService interface {
	CreatePurchase(ctx context.Context, purchase *dto.CreatePurchaseDTO) error
	GetPurchaseById(ctx context.Context, input *dto.GetPurchaseByIdDTO) (*dto.ProductResponse, error)
	GetAllPurchaseByUserId(ctx context.Context, input *dto.GetAllPurchaseByUserIdDTO) ([]*dto.PurchaseResponseDTO, error)
	GetAllPurchases(ctx context.Context, input *dto.GetAllPurchasesDTO) ([]*dto.PurchaseResponseDTO, error)
}
type RatingService interface {
	CreateRating(ctx context.Context, rating *dto.CreateRatingDTO) error
	UpdateRating(ctx context.Context, updateIt *dto.UpdateRatingDTO) error
	DeleteRating(ctx context.Context, deleteIt *dto.DeleteRatingDTO) error
	GetRatingById(ctx context.Context, input *dto.GetRatingByIdDTO) (*dto.RatingResponseDTO, error)
	GetRatingByUserId(ctx context.Context, input *dto.GetRatingByUserIdDTO) ([]*dto.RatingResponseDTO, error)
	GetAllByProductId(ctx context.Context, input *dto.GetAllByProductIdDTO) ([]*dto.RatingResponseDTO, error)
}
