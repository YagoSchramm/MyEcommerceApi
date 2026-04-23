package usecase_test

import (
	"context"
	"testing"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
	"github.com/google/uuid"
)

func buildProductTest(t *testing.T) (*usecase.ProductUsecase, uuid.UUID, string, func()) {
	t.Helper()

	db := usecase.OpenTestDB(t)
	_, userID, username := usecase.CreateTestUser(t, db, "Yago", "product-user-", []entity.Role{entity.RoleAdmin})

	productRepo := repository.NewProductRepository(db)
	productSrv := usecase.NewProductUsecase(productRepo)
	cleanup := func() {
		_ = db.Close()
	}

	return productSrv, userID, username, cleanup
}

func TestProductUsecase(t *testing.T) {
	usc, userID, username, cleanup := buildProductTest(t)
	defer cleanup()

	var productID uuid.UUID

	t.Run("Create Product", func(t *testing.T) {
		ctx := context.TODO()
		productMock := &dto.CreateProductDTO{
			UserID:      userID,
			UserName:    username,
			Name:        "Vaso de planta",
			Value:       38.99,
			Image:       "example-product.jpg",
			Stock:       54,
			Description: "Um vaso de planta com 30 cm de altura e 10 cm de diametro",
		}
		productID, err := usc.CreateProduct(ctx, productMock)
		if err != nil {
			t.Fatalf("Erro na criacao do produto: %s", err)
		}
		t.Log(productID)
	})

	t.Run("Get Product By Id", func(t *testing.T) {
		ctx := context.TODO()
		input := &dto.GetProductByIdDTO{
			ID: productID,
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
			ID: productID,
		}
		product, err := usc.GetAllProducts(ctx, input)
		if err != nil {
			t.Fatalf("Erro na busca dos produtos: %s", err)
		}
		t.Log(product)
	})
}
