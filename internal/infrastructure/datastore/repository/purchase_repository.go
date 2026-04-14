package repository

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
)

type PurchaseRepository struct {
	db *sql.DB
}

func NewPurchaseRepository(db *sql.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

//go:embed _query/purchase/create_purchase.sql
var createPurchaseQuery string

//go:embed _query/purchase/findByUser_purchase.sql
var findPurchaseByUserQuery string

//go:embed _query/purchase/findById_purchase.sql
var findPurchaseByIdQuery string

//go:embed _query/purchase/getAll_purchase.sql
var getAllPurchaseQuery string

func (p *PurchaseRepository) CreatePurchase(ctx context.Context, input entity.Purchase) error {
	_, err := p.db.ExecContext(ctx, createPurchaseQuery,
		input.ID,
		input.ProductID,
		input.UserID,
		input.Value,
		input.Quantity,
		input.CreatedAt,
	)

	return err
}
func (p *PurchaseRepository) GetPurchaseById(ctx context.Context, id string) (*entity.Purchase, error) {
	rows, err := p.db.QueryContext(ctx, findPurchaseByIdQuery, id)
	if err != nil {
		return nil, err
	}
	var purchase entity.Purchase
	err = rows.Scan(
		purchase.ID,
		purchase.ProductID,
		purchase.UserID,
		purchase.Value,
		purchase.Quantity,
		purchase.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &purchase, nil
}
func (p *PurchaseRepository) GetAllPurchaseByUserId(ctx context.Context, user_id string) ([]*entity.Purchase, error) {
	rows, err := p.db.QueryContext(ctx, findPurchaseByUserQuery, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var purchaseList []*entity.Purchase
	for rows.Next() {
		var purchase entity.Purchase
		err = rows.Scan(
			purchase.ID,
			purchase.ProductID,
			purchase.UserID,
			purchase.Value,
			purchase.Quantity,
			purchase.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		purchaseList = append(purchaseList, &purchase)
	}
	return purchaseList, nil
}
func (p *PurchaseRepository) GetPriceByProductId(ctx context.Context, product_id string) (float32, error) {
	var price float32
	err := p.db.QueryRowContext(ctx, "SELECT value FROM products WHERE product_id = $1", product_id).Scan(&price)
	if err != nil {
		return 0, err
	}
	return price, nil
}
func (p *PurchaseRepository) GetAllPurchases(ctx context.Context) ([]*entity.Purchase, error) {
	rows, err := p.db.QueryContext(ctx, getAllPurchaseQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var purchaseList []*entity.Purchase
	for rows.Next() {
		var purchase entity.Purchase
		err = rows.Scan(
			purchase.ID,
			purchase.ProductID,
			purchase.UserID,
			purchase.Value,
			purchase.Quantity,
			purchase.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		purchaseList = append(purchaseList, &purchase)
	}
	return purchaseList, nil
}
