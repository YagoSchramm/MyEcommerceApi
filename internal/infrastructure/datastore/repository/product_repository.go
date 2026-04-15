package repository

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"time"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service/dto"
	"github.com/google/uuid"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func NewProuductRepository(db *sql.DB) *ProductRepository {
	return NewProductRepository(db)
}

//go:embed _query/product/create_product.sql
var createProductQuery string

//go:embed _query/product/delete_product.sql
var deleteProductQuery string

//go:embed _query/rating/avg_rating.sql
var getAvgRating string

//go:embed _query/product/getAll_product.sql
var getAllProductQuery string

//go:embed _query/product/getById_product.sql
var getProductByIdQuery string

//go:embed _query/product/update_product.sql
var updateProductQuery string

func (pr *ProductRepository) CreateProduct(ctx context.Context, input entity.Product) (*uuid.UUID, error) {
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	_, err = tx.ExecContext(
		ctx,
		createProductQuery,
		input.ID,
		input.UserID,
		input.UserName,
		input.Name,
		input.Value,
		input.Image,
		input.Stock,
		input.Description,
		input.CreatedAt,
		input.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &input.ID, nil
}
func (pr *ProductRepository) UpdateProduct(ctx context.Context, updateIt dto.UpdateProductDTO) error {
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	lockQuery := `
		SELECT id FROM products
		WHERE id = $1 AND deleted_at IS NULL
		FOR UPDATE
	`

	var id string
	err = tx.QueryRowContext(ctx, lockQuery, updateIt.ID).Scan(&id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		updateProductQuery,
		updateIt.Name,
		updateIt.Value,
		updateIt.Image,
		updateIt.Stock,
		updateIt.Description,
		time.Now(),
		updateIt.ID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
func (pr *ProductRepository) DeleteProduct(ctx context.Context, deleteIt dto.DeleteProductDTO) error {
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, deleteProductQuery, deleteIt.ID, deleteIt.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		tx.Rollback()
		return fmt.Errorf("product not found or unauthorized")
	}

	return tx.Commit()

}
func (pr *ProductRepository) GetProductById(ctx context.Context, id string) (*entity.Product, error) {
	var p entity.Product

	err := pr.db.QueryRowContext(ctx, getProductByIdQuery, id).Scan(
		&p.ID,
		&p.UserID,
		&p.UserName,
		&p.Name,
		&p.Value,
		&p.Image,
		&p.AvgRating,
		&p.TotalRatings,
		&p.Stock,
		&p.Description,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	err = pr.db.QueryRowContext(ctx, getAvgRating, id).Scan(
		&p.AvgRating,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (pr *ProductRepository) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	rows, err := pr.db.QueryContext(ctx, getAllProductQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var p entity.Product

		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.UserName,
			&p.Name,
			&p.Value,
			&p.Image,
			&p.Stock,
			&p.AvgRating,
			&p.TotalRatings,
			&p.Description,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		err = pr.db.QueryRowContext(ctx, getAvgRating, p.ID).Scan(
			&p.AvgRating,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return products, nil
}
