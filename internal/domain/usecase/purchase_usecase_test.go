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

func buildPurchaseTest(t *testing.T) (*usecase.PurchaseUsecase, uuid.UUID, uuid.UUID, func()) {
	t.Helper()

	db := usecase.OpenTestDB(t)
	_, userID, username := usecase.CreateTestUser(t, db, "Yago", "purchase-user-", []entity.Role{entity.RoleAdmin})

	productRepo := repository.NewProductRepository(db)
	productSrv := usecase.NewProductUsecase(productRepo)
	productMock := &dto.CreateProductDTO{
		UserID:      userID,
		UserName:    username,
		Name:        "Vaso de planta",
		Value:       38.99,
		Image:       "example-purchase.jpg",
		Stock:       54,
		Description: "Um vaso de planta com 30 cm de altura e 10 cm de diametro",
	}
	productID, err := productSrv.CreateProduct(context.Background(), productMock)
	if err != nil {
		_ = db.Close()
		t.Fatalf("falha ao criar produto de teste: %v", err)
	}

	purchaseRepo := repository.NewPurchaseRepository(db)
	purchaseUsc := usecase.NewPurchaseUsecase(purchaseRepo)
	cleanup := func() {
		_ = db.Close()
	}

	return purchaseUsc, userID, *productID, cleanup
}

func TestPurchase(t *testing.T) {
	usc, userID, productID, cleanup := buildPurchaseTest(t)
	defer cleanup()

	var purchaseID *uuid.UUID

	t.Run("Create Purchase", func(t *testing.T) {
		ctx := context.TODO()
		purchase := &dto.CreatePurchaseDTO{
			ProductID: productID,
			UserID:    userID,
			Quantity:  2,
		}
		var err error
		purchaseID, err = usc.CreatePurchase(ctx, purchase)
		if err != nil {
			t.Fatalf("Erro ao criar compra: %s", err)
		}
	})

	t.Run("Get Purchase By Id", func(t *testing.T) {
		ctx := context.TODO()
		input := &dto.GetPurchaseByIdDTO{
			ID: *purchaseID,
		}
		purchases, err := usc.GetPurchaseById(ctx, input)
		if err != nil {
			t.Fatalf("Erro ao buscar compra: %s", err)
		}
		t.Log(purchases)
	})

	t.Run("Get all Purchases By User Id", func(t *testing.T) {
		ctx := context.TODO()
		input := &dto.GetAllPurchaseByUserIdDTO{
			UserID: userID,
		}
		purchases, err := usc.GetAllPurchaseByUserId(ctx, input)
		if err != nil {
			t.Fatalf("Erro ao buscar todas as compras do usuario: %s", err)
		}
		t.Log(purchases)
	})

	t.Run("Get all Purchases", func(t *testing.T) {
		ctx := context.TODO()
		input := &dto.GetAllPurchasesDTO{
			ID: uuid.New(),
		}
		purchases, err := usc.GetAllPurchases(ctx, input)
		if err != nil {
			t.Fatalf("Erro ao buscar todas as compras: %s", err)
		}
		t.Log(purchases)
	})
}
