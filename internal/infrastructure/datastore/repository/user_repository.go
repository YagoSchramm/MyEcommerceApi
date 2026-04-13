package repository

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"time"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service/dto"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

//go:embed _query/user/create_user.sql
var createUserQuery string

//go:embed _query/user/delete_user.sql
var deleteUserQuery string

//go:embed _query/user/getAll_user.sql
var getAllUserQuery string

//go:embed _query/user/getById_user.sql
var getUserByIdQuery string

//go:embed _query/user/getByRole_user.sql
var getUserByRoleQuery string

//go:embed _query/user/update_user.sql
var updateUserQuery string

func (ur *UserRepository) CreateUser(ctx context.Context, input entity.User) error {
	tx, err := ur.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(
		ctx,
		createUserQuery,
		input.ID,
		input.Name,
		input.Email,
		input.Password,
		input.Roles,
		input.CreatedAt,
		input.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (ur *UserRepository) UpdateUser(ctx context.Context, updateIt dto.UpdateUserDTO) error {
	tx, err := ur.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	lockQuery := `
		SELECT id FROM users
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
		updateUserQuery,
		updateIt.Name,
		updateIt.Roles,
		time.Now(),
		updateIt.ID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (ur *UserRepository) DeleteUser(ctx context.Context, deleteIt string) error {
	tx, err := ur.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, deleteUserQuery, time.Now(), deleteIt)
	if err != nil {
		tx.Rollback()
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		tx.Rollback()
		return fmt.Errorf("user not found")
	}

	return tx.Commit()
}

func (ur *UserRepository) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	var u entity.User

	err := ur.db.QueryRowContext(ctx, getUserByIdQuery, id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Roles,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (ur *UserRepository) GetUserByRole(ctx context.Context, role entity.Role) ([]*entity.User, error) {
	rows, err := ur.db.QueryContext(ctx, getUserByRoleQuery, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User

	for rows.Next() {
		var u entity.User

		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Password,
			&u.Roles,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil
}

func (ur *UserRepository) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	rows, err := ur.db.QueryContext(ctx, getAllUserQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User

	for rows.Next() {
		var u entity.User

		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Password,
			&u.Roles,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil
}
