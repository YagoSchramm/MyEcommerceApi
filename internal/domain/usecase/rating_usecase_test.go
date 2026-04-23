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

func buildRatingTest(t *testing.T) (*usecase.RatingUsecase, uuid.UUID, uuid.UUID, uuid.UUID, func()) {
	t.Helper()

	db := usecase.OpenTestDB(t)
	_, userID, _ := usecase.CreateTestUser(t, db, "Rating Test User", "rating-user-", []entity.Role{entity.RoleBuyer, entity.RoleSeller})

	productRepo := repository.NewProductRepository(db)
	productSrv := usecase.NewProductUsecase(productRepo)
	productDTO := &dto.CreateProductDTO{
		UserID:      userID,
		UserName:    "Rating Test User",
		Name:        "Rating Test Product " + uuid.NewString(),
		Value:       99.99,
		Image:       "rating-test.jpg",
		Stock:       100,
		Description: "A test product for rating integration testing with proper description length requirements.",
	}
	productIDPtr, err := productSrv.CreateProduct(context.Background(), productDTO)
	if err != nil {
		_ = db.Close()
		t.Fatalf("falha ao criar produto: %v", err)
	}
	if productIDPtr == nil {
		_ = db.Close()
		t.Fatal("falha ao criar produto: id retornado nulo")
	}

	var productID uuid.UUID
	row := db.QueryRowContext(context.Background(), "SELECT id FROM products WHERE name = $1 LIMIT 1", productDTO.Name)
	if err := row.Scan(&productID); err != nil {
		_ = db.Close()
		t.Fatalf("falha ao localizar produto criado: %v", err)
	}

	purchaseRepo := repository.NewPurchaseRepository(db)
	purchaseSrv := usecase.NewPurchaseUsecase(purchaseRepo)
	purchaseDTO := &dto.CreatePurchaseDTO{
		ProductID: productID,
		UserID:    userID,
		Quantity:  2,
	}
	purchaseID, purchaseErr := purchaseSrv.CreatePurchase(context.Background(), purchaseDTO)
	if purchaseErr != nil {
		_ = db.Close()
		t.Fatalf("falha ao criar compra: %v", purchaseErr)
	}

	ratingRepo := repository.NewRatingRepository(db)
	ratingSrv := usecase.NewRatingUsecase(ratingRepo)

	cleanup := func() {
		db.Close()
	}

	return ratingSrv, userID, productID, *purchaseID, cleanup
}

func TestRatingUsecase(t *testing.T) {
	usc, userID, productID, purchaseID, cleanup := buildRatingTest(t)
	defer cleanup()

	ctx := context.Background()
	var ratingID uuid.UUID

	t.Run("CreateRating", func(t *testing.T) {
		createDTO := &dto.CreateRatingDTO{
			UserID:      userID,
			UserName:    "Rating Test User",
			ProdutctID:  productID,
			PurchaseID:  purchaseID,
			Rating:      4.5,
			Description: "Excelente produto, muito bom!",
		}
		err := usc.CreateRating(ctx, createDTO)
		if err != nil {
			t.Fatalf("CreateRating falhou: %v", err)
		}
		t.Logf("Avaliacao criada com sucesso para compra %s", purchaseID)
	})

	t.Run("GetRatingByUserId", func(t *testing.T) {
		input := &dto.GetRatingByUserIdDTO{UserID: userID}
		ratings, err := usc.GetRatingByUserId(ctx, input)
		if err != nil {
			t.Fatalf("GetRatingByUserId falhou: %v", err)
		}
		if len(ratings) == 0 {
			t.Fatalf("esperado pelo menos uma avaliacao do usuario")
		}
		ratingID = ratings[0].ID
	})

	t.Run("GetAllByProductId", func(t *testing.T) {
		input := &dto.GetAllRatingByProductIdDTO{ProductID: productID}
		ratings, err := usc.GetAllByProductId(ctx, input)
		if err != nil {
			t.Fatalf("GetAllByProductId falhou: %v", err)
		}
		if len(ratings) == 0 {
			t.Fatalf("esperado pelo menos uma avaliacao para o produto")
		}
	})

	t.Run("GetRatingById", func(t *testing.T) {
		input := &dto.GetRatingByIdDTO{ID: ratingID}
		rating, err := usc.GetRatingById(ctx, input)
		if err != nil {
			t.Fatalf("GetRatingById falhou: %v", err)
		}
		if rating == nil || rating.ID != ratingID {
			t.Fatalf("esperado id de avaliacao %v, obtido %v", ratingID, rating)
		}
	})

	t.Run("UpdateRating", func(t *testing.T) {
		updateDTO := &dto.UpdateRatingDTO{
			ID:     ratingID,
			Rating: 5.0,
		}
		err := usc.UpdateRating(ctx, updateDTO)
		if err != nil {
			t.Fatalf("UpdateRating falhou: %v", err)
		}

		updated, err := usc.GetRatingById(ctx, &dto.GetRatingByIdDTO{ID: ratingID})
		if err != nil {
			t.Fatalf("GetRatingById apos atualizacao falhou: %v", err)
		}
		if updated.Rating != 5.0 {
			t.Fatalf("esperado avaliacao 5.0, obtido %v", updated.Rating)
		}
	})

	t.Run("DeleteRating", func(t *testing.T) {
		deleteDTO := &dto.DeleteRatingDTO{
			ID:         ratingID,
			UserID:     userID,
			ProdutctID: productID,
		}
		err := usc.DeletRating(ctx, deleteDTO)
		if err != nil {
			t.Fatalf("DeleteRating falhou: %v", err)
		}

		_, err = usc.GetRatingById(ctx, &dto.GetRatingByIdDTO{ID: ratingID})
		if err == nil {
			t.Fatal("esperado GetRatingById falhar apos deletar")
		}
	})
}
