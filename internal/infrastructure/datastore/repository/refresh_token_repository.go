package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
)

type RefreshTokenRepository struct {
	db *sql.DB
}

func NewRefreshTokenRepository(db *sql.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

// Save stores a new refresh token in the database
func (r *RefreshTokenRepository) Save(userID, token string) error {
	query := `
		INSERT INTO refresh_tokens (user_id, token, expiry)
		VALUES ($1, $2, $3)
	`

	expiry := time.Now().Add(7 * 24 * time.Hour)

	_, err := r.db.Exec(query, userID, token, expiry)
	if err != nil {
		return err
	}

	return nil
}

// Exists checks if a refresh token exists in the database
func (r *RefreshTokenRepository) Exists(userID, token string) bool {
	query := `
		SELECT token FROM refresh_tokens 
		WHERE user_id = $1 AND token = $2 AND expiry > NOW()
	`

	var existingToken string
	err := r.db.QueryRow(query, userID, token).Scan(&existingToken)

	if err != nil && err != sql.ErrNoRows {
		return false
	}

	return err != sql.ErrNoRows
}

// Delete removes a refresh token from the database
func (r *RefreshTokenRepository) Delete(userID, token string) error {
	query := `
		DELETE FROM refresh_tokens 
		WHERE user_id = $1 AND token = $2
	`

	result, err := r.db.Exec(query, userID, token)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("refresh token not found")
	}

	return nil
}

// GetByUserID retrieves all refresh tokens for a user
func (r *RefreshTokenRepository) GetByUserID(userID string) ([]entity.RefreshToken, error) {
	query := `
		SELECT user_id, token, expiry FROM refresh_tokens 
		WHERE user_id = $1 AND expiry > NOW()
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []entity.RefreshToken

	for rows.Next() {
		var rt entity.RefreshToken
		err := rows.Scan(&rt.UserID, &rt.Token, &rt.Expiry)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, rt)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tokens, nil
}

// DeleteAllByUserID removes all refresh tokens for a user (useful for logout)
func (r *RefreshTokenRepository) DeleteAllByUserID(userID string) error {
	query := `
		DELETE FROM refresh_tokens WHERE user_id = $1
	`

	_, err := r.db.Exec(query, userID)
	return err
}
