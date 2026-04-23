package domain

import (
	"context"
	"mime/multipart"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
)

type DomainUsecase interface {
	UserUsecase
	ProductUsecase
	PurchaseUsecase
	RatingUsecase
}
type UserUsecase interface {
	CreateUser(ctx context.Context, user *dto.CreateUserDTO) error
	UpdateUser(ctx context.Context, updateIt *dto.UpdateUserDTO) error
	DeleteUser(ctx context.Context, deleteIt *dto.DeleteUserDTO) error
	GetUserById(ctx context.Context, input *dto.GetUserByIdDTO) (*dto.UserResponseDTO, error)
	GetUserByRole(ctx context.Context, input *dto.GetUserByRoleDTO) ([]*dto.UserResponseDTO, error)
	GetAllUsers(ctx context.Context, input *dto.GetAllUsersDTO) ([]*dto.UserResponseDTO, error)
}
type ProductUsecase interface {
	CreateProduct(ctx context.Context, product *dto.CreateProductDTO) error
	UpdateProduct(ctx context.Context, updateIt *dto.UpdateProductDTO) error
	DeleteProduct(ctx context.Context, deleteIt *dto.DeleteProductDTO) error
	GetProductById(ctx context.Context, input *dto.GetProductByIdDTO) (*dto.ProductResponse, error)
	GetAllProducts(ctx context.Context, input *dto.GetAllProductsDTO) ([]*dto.ProductResponse, error)
}
type PurchaseUsecase interface {
	CreatePurchase(ctx context.Context, purchase *dto.CreatePurchaseDTO) error
	GetPurchaseById(ctx context.Context, input *dto.GetPurchaseByIdDTO) (*dto.ProductResponse, error)
	GetAllPurchaseByUserId(ctx context.Context, input *dto.GetAllPurchaseByUserIdDTO) ([]*dto.PurchaseResponseDTO, error)
	GetAllPurchases(ctx context.Context, input *dto.GetAllPurchasesDTO) ([]*dto.PurchaseResponseDTO, error)
}
type RatingUsecase interface {
	CreateRating(ctx context.Context, rating *dto.CreateRatingDTO) error
	UpdateRating(ctx context.Context, updateIt *dto.UpdateRatingDTO) error
	DeleteRating(ctx context.Context, deleteIt *dto.DeleteRatingDTO) error
	GetRatingById(ctx context.Context, input *dto.GetRatingByIdDTO) (*dto.RatingResponseDTO, error)
	GetRatingByUserId(ctx context.Context, input *dto.GetRatingByUserIdDTO) ([]*dto.RatingResponseDTO, error)
	GetAllByProductId(ctx context.Context, input *dto.GetAllRatingByProductIdDTO) ([]*dto.RatingResponseDTO, error)
}
type ImageUsecase interface {
	Save(file multipart.File, filename string) (*dto.ImageResponseDTO, error)
}
