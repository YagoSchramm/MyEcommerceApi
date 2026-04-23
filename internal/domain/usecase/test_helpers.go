package usecase

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/YagoSchramm/myecommerce-api/internal/foundation"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
	"github.com/google/uuid"
)

func OpenTestDB(t *testing.T) *sql.DB {
	t.Helper()

	conn := os.Getenv("TEST_DB_URL")
	if conn == "" {
		conn = os.Getenv("DB_URL")
	}
	if conn == "" {
		conn = "postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable"
	}

	db, err := foundation.NewPostgresDB(conn)
	if err != nil {
		t.Skipf("Skipping integration test because DB connection failed: %v", err)
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		t.Skipf("Skipping integration test because DB is unavailable: %v", err)
	}

	return db
}

func NewTestTokenService() *service.TokenService {
	return service.NewTokenService(os.Getenv("JWT_SECRET"))
}

func CreateTestUser(t *testing.T, db *sql.DB, name, emailPrefix string, roles []entity.Role) (*UserUsecase, uuid.UUID, string) {
	t.Helper()

	userRepo := repository.NewUserRepository(db)
	userUsc := NewUserUsecase(userRepo, NewTestTokenService())
	email := emailPrefix + uuid.NewString() + "@example.com"

	createDTO := &dto.CreateUserDTO{
		Name:     name,
		Email:    email,
		Password: "password123",
		Roles:    roles,
	}

	if err := userUsc.CreateUser(context.Background(), createDTO); err != nil {
		t.Fatalf("falha ao criar usuário de teste: %v", err)
	}

	var userID uuid.UUID
	row := db.QueryRowContext(context.Background(), "SELECT id FROM users WHERE email = $1 LIMIT 1", email)
	if err := row.Scan(&userID); err != nil {
		t.Fatalf("falha ao localizar usuário criado: %v", err)
	}

	return userUsc, userID, email
}
