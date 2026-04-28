package domain

import (
	"context"
	"mime/multipart"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/google/uuid"
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
	GetPurchaseById(ctx context.Context, input *dto.GetPurchaseByIdDTO) (*dto.PurchaseResponseDTO, error)
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

type TokenService interface {
	GenerateTokens(userID string, roles []string) (string, string, error)
	ValidateRefreshToken(tokenStr string) (*service.RefreshClaims, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, input entity.User) error
	UpdateUser(ctx context.Context, updateIt dto.UpdateUserDTO) error
	DeleteUser(ctx context.Context, deleteIt dto.DeleteUserDTO) error
	GetUserById(ctx context.Context, id string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByRole(ctx context.Context, role entity.Role) ([]*entity.User, error)
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, input entity.Product) (*uuid.UUID, error)
	UpdateProduct(ctx context.Context, updateIt dto.UpdateProductDTO) error
	DeleteProduct(ctx context.Context, deleteIt dto.DeleteProductDTO) error
	GetProductById(ctx context.Context, id string) (*entity.Product, error)
	GetAllProducts(ctx context.Context) ([]*entity.Product, error)
}

type PurchaseRepository interface {
	CreatePurchase(ctx context.Context, input entity.Purchase) (*uuid.UUID, error)
	GetPurchaseById(ctx context.Context, id string) (*entity.Purchase, error)
	GetAllPurchaseByUserId(ctx context.Context, userID string) ([]*entity.Purchase, error)
	GetPriceByProductId(ctx context.Context, productID string) (float32, error)
	GetAllPurchases(ctx context.Context) ([]*entity.Purchase, error)
}

type RatingRepository interface {
	CreateRating(ctx context.Context, input entity.Rating) error
	UpdateRating(ctx context.Context, updateIt dto.UpdateRatingDTO) error
	DeleteRating(ctx context.Context, deleteIt *dto.DeleteRatingDTO) error
	GetRatingById(ctx context.Context, id string) (*entity.Rating, error)
	GetRatingByUserId(ctx context.Context, userID string) ([]*entity.Rating, error)
	GetAllByProductId(ctx context.Context, productID string) ([]*entity.Rating, error)
}

type ImageRepository interface {
	Save(file multipart.File, filename string) (string, error)
}

type RefreshTokenRepository interface {
	Save(userID, token string) error
	Exists(userID, token string) bool
	Delete(userID, token string) error
}

type RefreshUserRepository interface {
	GetByID(userID string) (*entity.User, error)
}
