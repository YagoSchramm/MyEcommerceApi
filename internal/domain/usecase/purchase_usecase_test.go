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

func buildPurchaseTest(t *testing.T) (*usecase.PurchaseUsecase, uuid.UUID, uuid.UUID) {
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
	productMock := &dto.CreateProductDTO{
		UserID:      user[0].ID,
		UserName:    user[0].Name,
		Name:        "Vaso de planta",
		Value:       38.99,
		Image:       "google.com/example",
		Stock:       54,
		Description: "Um vaso de planta com 30 cm de altura e 10 cm de diâmetro",
	}
	productId, _ := productSrv.CreateProduct(context.TODO(), productMock)
	purchaseRepo := repository.NewPurchaseRepository(db)
	PurchaseUsc := usecase.NewPurchaseUsecase(purchaseRepo)
	return PurchaseUsc, user[0].ID, *productId
}
func TestPurchase(t *testing.T) {
	usc, user_id, product_id := buildPurchaseTest(t)
	var purchase_id *uuid.UUID
	t.Run("Create Purchase", func(t *testing.T) {
		ctx := context.TODO()
		purchase := &dto.CreatePurchaseDTO{
			ProductID: product_id,
			UserID:    user_id,
			Quantity:  2,
		}
		var err error
		purchase_id, err = usc.CreatePurchase(ctx, purchase)
		if err != nil {
			t.Fatalf("Erro ao criar compra: %s", err)
		}
	})
	t.Run("Get Purchase By Id", func(t *testing.T) {
		ctx := context.TODO()
		input := &dto.GetPurchaseByIdDTO{
			ID: *purchase_id,
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
			UserID: user_id,
		}
		purchases, err := usc.GetAllPurchaseByUserId(ctx, input)
		if err != nil {
			t.Fatalf("Erro ao buscar todas as compras do usuário: %s", err)

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
