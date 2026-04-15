package datastore

import (
	"context"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
)

type Store interface {
	UserRepository
	ProductRepository
	RatingRepository
	PurchaseRepository
}
type UserRepository interface {
	CreateUser(ctx context.Context, input entity.User) error
	UpdateUser(ctx context.Context, updateIt dto.UpdateUserDTO) error
	DeleteUser(ctx context.Context, deleteIt string) error
	GetUserById(ctx context.Context, id string) (*entity.User, error)
	GetUserByRole(ctx context.Context, role entity.Role) ([]*entity.User, error)
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
}
type ProductRepository interface {
	CreateProduct(ctx context.Context, input entity.Product) error
	UpdateProduct(ctx context.Context, updateIt dto.UpdateProductDTO) error
	DeleteProduct(ctx context.Context, deleteIt string) error
	GetProductById(ctx context.Context, id string) (*entity.Product, error)
	GetAllProducts(ctx context.Context) ([]*entity.Product, error)
}
type RatingRepository interface {
	CreateRating(ctx context.Context, input entity.Rating) error
	UpdateRating(ctx context.Context, updateIt dto.UpdateProductDTO) error
	DeleteRating(ctx context.Context, deleteIt string) error
	GetRatingById(ctx context.Context, id string) (*entity.Rating, error)
	GetRatingByUserId(ctx context.Context, user_id string) ([]*entity.Rating, error)
	GetAllByProductId(ctx context.Context, productId string) ([]*entity.Rating, error)
}
type PurchaseRepository interface {
	CreatePurchase(ctx context.Context, input entity.Purchase) error
	GetPurchaseById(ctx context.Context, id string) (*entity.Purchase, error)
	GetAllPurchaseByUserId(ctx context.Context, user_id string) ([]*entity.Purchase, error)
	GetAllPurchases(ctx context.Context) ([]*entity.Purchase, error)
}
