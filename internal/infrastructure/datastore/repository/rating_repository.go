package repository

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"time"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
)

type RatingRepository struct {
	db *sql.DB
}

func NewRatingRepository(db *sql.DB) *RatingRepository {
	return &RatingRepository{db: db}
}

//go:embed _query/rating/create_rating.sql
var createRatingQuery string

//go:embed _query/rating/update_rating.sql
var updateRatingQuery string

//go:embed _query/rating/delete_rating.sql
var deleteRatingQuery string

//go:embed _query/rating/getById_rating.sql
var getRatingByIdQuery string

//go:embed _query/rating/getByProductId_rating.sql
var getRatingByProductIdQuery string

//go:embed _query/rating/getByuser_rating.sql
var getRatingByUserQuery string

//go:embed _query/rating/validate_rating.sql
var validateRatingQuery string

//go:embed _query/rating/checkDuplicate_rating.sql
var checkDuplicateQuery string

//go:embed _query/product/update_ratingOnProduct.sql
var updateProductRatingQuery string

func (r *RatingRepository) CreateRating(ctx context.Context, input entity.Rating) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var purchaseID string
	err = tx.QueryRowContext(ctx, validateRatingQuery, input.PurchaseID, input.UserID).Scan(&purchaseID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("invalid purchase or unauthorized")
	}

	var existingID string
	err = tx.QueryRowContext(ctx, checkDuplicateQuery, input.PurchaseID).Scan(&existingID)
	if err == nil {
		tx.Rollback()
		return fmt.Errorf("rating already exists for this purchase")
	}
	_, err = tx.ExecContext(ctx, updateProductRatingQuery, input.ProductID)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(
		ctx,
		createRatingQuery,
		input.ID,
		input.UserID,
		input.UserName,
		input.PurchaseID,
		input.ProductID,
		input.Rating,
		input.Description,
		input.CreatedAt,
		input.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
func (r *RatingRepository) UpdateRating(ctx context.Context, updateIt dto.UpdateRatingDTO) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		updateRatingQuery,
		updateIt.Rating,
		time.Now(),
		updateIt.ID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
func (r *RatingRepository) DeleteRating(ctx context.Context, deletIt *dto.DeleteRatingDTO) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	res, err := tx.ExecContext(ctx, deleteRatingQuery, deletIt.ID, deletIt.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, updateProductRatingQuery, deletIt.ProdutctID)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		tx.Rollback()
		return fmt.Errorf("not found or unauthorized")
	}

	return tx.Commit()
}
func (r *RatingRepository) GetRatingById(ctx context.Context, id string) (*entity.Rating, error) {

	var rating entity.Rating

	err := r.db.QueryRowContext(ctx, getRatingByIdQuery, id).Scan(
		&rating.ID,
		&rating.UserID,
		&rating.UserName,
		&rating.PurchaseID,
		&rating.ProductID,
		&rating.Rating,
		&rating.Description,
		&rating.CreatedAt,
		&rating.UpdatedAt,
		&rating.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &rating, nil
}
func (r *RatingRepository) GetRatingByUserId(ctx context.Context, userID string) ([]*entity.Rating, error) {

	rows, err := r.db.QueryContext(ctx, getRatingByUserQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []*entity.Rating

	for rows.Next() {
		var rating entity.Rating

		err := rows.Scan(
			&rating.ID,
			&rating.UserID,
			&rating.UserName,
			&rating.PurchaseID,
			&rating.ProductID,
			&rating.Rating,
			&rating.Description,
			&rating.CreatedAt,
			&rating.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		ratings = append(ratings, &rating)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}
func (r *RatingRepository) GetAllByProductId(ctx context.Context, productID string) ([]*entity.Rating, error) {

	rows, err := r.db.QueryContext(ctx, getRatingByProductIdQuery, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []*entity.Rating

	for rows.Next() {
		var rating entity.Rating

		err := rows.Scan(
			&rating.ID,
			&rating.UserID,
			&rating.UserName,
			&rating.PurchaseID,
			&rating.ProductID,
			&rating.Rating,
			&rating.Description,
			&rating.CreatedAt,
			&rating.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		ratings = append(ratings, &rating)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}
