package usecase_test

import (
	"context"
	"os"
	"testing"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/YagoSchramm/myecommerce-api/internal/foundation"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
	"github.com/google/uuid"
)

func buildProductTest(t *testing.T) (*usecase.ProductUsecase, uuid.UUID, string) {
	t.Helper()
	conn := "postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable"

	db, _ := foundation.NewPostgresDB(conn)
	userRepo := repository.NewUserRepository(db)
	secret := os.Getenv("JWT-SECRET")
	jwtSrv := service.NewTokenService(secret)
	userUsc := usecase.NewUserUsecase(userRepo, jwtSrv)
	userMock := dto.CreateUserDTO{
		Name:     "Yago",
		Email:    "yago@gmail.com",
		Password: "12345",
		Roles:    []entity.Role{"admin"},
	}
	userUsc.CreateUser(context.TODO(), &userMock)

	user, _ := userUsc.GetUserByRole(context.TODO(), &dto.GetUserByRoleDTO{Role: "admin"})
	productRepo := repository.NewProductRepository(db)
	productSrv := usecase.NewProductUsecase(productRepo)
	return productSrv, user[0].ID, user[0].Name
}
func TestProductUsecase(t *testing.T) {
	usc, user_id, username := buildProductTest(t)
	var productId uuid.UUID
	t.Run("Create Product", func(t *testing.T) {
		ctx := context.TODO()
		productMock := &dto.CreateProductDTO{
			UserID:      user_id,
			UserName:    username,
			Name:        "Vaso de planta",
			Value:       38.99,
			Image:       "google.com/example",
			Stock:       54,
			Description: "Um vaso de planta com 30 cm de altura e 10 cm de diâmetro",
		}
		productId, err := usc.CreateProduct(ctx, productMock)
		if err != nil {
			t.Fatalf("Erro na criação do produto: %s", err)
		}
		t.Log(productId)

	})
	t.Run("Get Product By Id", func(t *testing.T) {
		ctx := context.TODO()
		input := &dto.GetProductByIdDTO{
			ID: productId,
		}
		product, err := usc.GetProductById(ctx, input)
		if err != nil {
			t.Fatalf("Erro na busca do produto: %s", err)
		}
		t.Log(product)
	})
	t.Run("Get All Products", func(t *testing.T) {
		ctx := context.TODO()
		input := &dto.GetAllProductsDTO{
			ID: productId,
		}
		product, err := usc.GetAllProducts(ctx, input)
		if err != nil {
			t.Fatalf("Erro na busca dos produtos: %s", err)
		}
		t.Log(product)
	})
}
