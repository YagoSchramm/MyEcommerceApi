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

func buildProductTest(t *testing.T) (*usecase.ProductUsecase, uuid.UUID, string, func()) {
	t.Helper()
	conn := "postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable"

	db, err := foundation.NewPostgresDB(conn)
	if err != nil {
		t.Skipf("Skipping integration test because DB connection failed: %v", err)
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		t.Skipf("Skipping integration test because DB is unavailable: %v", err)
	}
	userRepo := repository.NewUserRepository(db)
	secret := os.Getenv("JWT_SECRET")
	jwtSrv := service.NewTokenService(secret)
	userUsc := usecase.NewUserUsecase(userRepo, jwtSrv)
	email := "product-user-" + uuid.NewString() + "@example.com"
	userMock := dto.CreateUserDTO{
		Name:     "Yago",
		Email:    email,
		Password: "password123",
		Roles:    []entity.Role{entity.RoleAdmin},
	}
	if err := userUsc.CreateUser(context.Background(), &userMock); err != nil {
		_ = db.Close()
		t.Fatalf("falha ao criar usuÃ¡rio de teste: %v", err)
	}

	user, err := userUsc.GetUserByRole(context.Background(), &dto.GetUserByRoleDTO{Role: string(entity.RoleAdmin)})
	if err != nil {
		_ = db.Close()
		t.Fatalf("falha ao buscar usuÃ¡rio por role: %v", err)
	}
	if len(user) == 0 {
		_ = db.Close()
		t.Fatal("esperado pelo menos um usuÃ¡rio admin no setup do teste")
	}
	productRepo := repository.NewProductRepository(db)
	productSrv := usecase.NewProductUsecase(productRepo)
	cleanup := func() {
		_ = db.Close()
	}
	return productSrv, user[0].ID, user[0].Name, cleanup
}
func TestProductUsecase(t *testing.T) {
	usc, userID, username, cleanup := buildProductTest(t)
	defer cleanup()
	var productId uuid.UUID
	t.Run("Create Product", func(t *testing.T) {
		ctx := context.TODO()
		productMock := &dto.CreateProductDTO{
			UserID:      userID,
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
