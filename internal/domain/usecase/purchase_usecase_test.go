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

func buildPurchaseTest(t *testing.T) (*usecase.PurchaseUsecase, uuid.UUID, uuid.UUID, func()) {
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

	email := "purchase-user-" + uuid.NewString() + "@example.com"
	userMock := dto.CreateUserDTO{
		Name:     "Yago",
		Email:    email,
		Password: "password123",
		Roles:    []entity.Role{entity.RoleAdmin},
	}
	if err := userUsc.CreateUser(context.Background(), &userMock); err != nil {
		_ = db.Close()
		t.Fatalf("falha ao criar usuário de teste: %v", err)
	}

	user, err := userUsc.GetUserByRole(context.Background(), &dto.GetUserByRoleDTO{Role: string(entity.RoleAdmin)})
	if err != nil {
		_ = db.Close()
		t.Fatalf("falha ao buscar usuário por role: %v", err)
	}
	if len(user) == 0 {
		_ = db.Close()
		t.Fatal("esperado pelo menos um usuário admin no setup do teste")
	}

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

	return purchaseUsc, user[0].ID, *productID, cleanup
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
