package domain

import (
	"context"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/service/dto"
)

type DomainService interface {
}
type UserService interface {
	CreateUser(ctx context.Context, user *dto.CreateUserDTO) error
	UpdateUser(ctx context.Context, updateIt *dto.UpdateUserDTO) error
	DeleteUser(ctx context.Context, deleteIt *dto.DeleteUserDTO) error
	GetUserById(ctx context.Context, input *dto.GetUserByIdDTO) (*dto.GetUserByIdDTO, error)
	GetUserByRole(ctx context.Context, input *dto.GetUserByRoleDTO) ([]*dto.GetUserByRoleDTO, error)
	GetAllUsers(ctx context.Context, input *dto.GetAllUsersDTO) ([]*dto.GetAllUsersDTO, error)
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
	GetPurchaseById(ctx context.Context, input *dto.GetPurchaseByIdDTO) (*dto.GetPurchaseByIdDTO, error)
	GetAllPurchaseByUserId(ctx context.Context, input *dto.GetAllPurchaseByUserIdDTO) ([]*dto.GetAllPurchaseByUserIdDTO, error)
	GetAllPurchases(ctx context.Context, input *dto.GetAllPurchasesDTO) ([]*dto.GetAllPurchasesDTO, error)
}
type RatingService interface {
	CreateRating(ctx context.Context, rating *dto.CreateRatingDTO) error
	UpdateRating(ctx context.Context, updateIt *dto.UpdateRatingDTO) error
	DeleteRating(ctx context.Context, deleteIt *dto.DeleteRatingDTO) error
	GetRatingById(ctx context.Context, input *dto.GetRatingByIdDTO) (*dto.GetRatingByIdDTO, error)
	GetRatingByUserId(ctx context.Context, input *dto.GetRatingByUserIdDTO) ([]*dto.GetRatingByUserIdDTO, error)
	GetAllByProductId(ctx context.Context, input *dto.GetAllByProductIdDTO) ([]*dto.GetAllByProductIdDTO, error)
}
