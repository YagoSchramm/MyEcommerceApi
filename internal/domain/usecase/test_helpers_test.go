package usecase_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/YagoSchramm/myecommerce-api/internal/foundation"
)

func openTestDB(t *testing.T) *sql.DB {
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
