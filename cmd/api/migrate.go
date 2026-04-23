package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/lib/pq"
)

func runMigrations(connStr string) error {
	if err := ensureDatabaseExists(connStr); err != nil {
		return err
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT NOW()
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Find migration files
	migrationDir := "internal/scripts/migrate"
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		return fmt.Errorf("failed to read migration directory: %w", err)
	}

	// Filter and sort migration files
	var migrations []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".up.sql") {
			migrations = append(migrations, file.Name())
		}
	}
	sort.Strings(migrations)

	// Execute migrations
	for _, migration := range migrations {
		version := strings.TrimSuffix(migration, ".up.sql")

		// Check if already applied
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM schema_migrations WHERE version = $1", version).Scan(&count)
		if err != nil {
			return fmt.Errorf("failed to check migration status: %w", err)
		}

		if count > 0 {
			fmt.Printf("Migration %s already applied, skipping\n", version)
			continue
		}

		// Read migration file
		content, err := os.ReadFile(filepath.Join(migrationDir, migration))
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", migration, err)
		}

		// Execute migration
		fmt.Printf("Applying migration %s...\n", version)
		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", migration, err)
		}

		// Record migration
		_, err = db.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", version)
		if err != nil {
			return fmt.Errorf("failed to record migration %s: %w", version, err)
		}

		fmt.Printf("Migration %s applied successfully\n", version)
	}

	fmt.Println("All migrations completed successfully")
	return nil
}

func ensureDatabaseExists(connStr string) error {
	dbName, maintenanceConn, err := maintenanceConnection(connStr)
	if err != nil {
		return err
	}

	adminDB, err := sql.Open("postgres", maintenanceConn)
	if err != nil {
		return fmt.Errorf("failed to connect to maintenance database: %w", err)
	}
	defer adminDB.Close()

	if err := adminDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping maintenance database: %w", err)
	}

	var exists bool
	if err := adminDB.QueryRow(`SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)`, dbName).Scan(&exists); err != nil {
		return fmt.Errorf("failed to check database existence: %w", err)
	}

	if exists {
		return nil
	}

	fmt.Printf("Database %s does not exist, creating it...\n", dbName)

	if _, err := adminDB.Exec(fmt.Sprintf("CREATE DATABASE %s", pq.QuoteIdentifier(dbName))); err != nil {
		return fmt.Errorf("failed to create database %q: %w", dbName, err)
	}

	return nil
}

func maintenanceConnection(connStr string) (string, string, error) {
	parsed, err := url.Parse(connStr)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse database url: %w", err)
	}

	dbName := strings.TrimPrefix(parsed.Path, "/")
	if dbName == "" {
		return "", "", fmt.Errorf("database name is missing in connection string")
	}

	parsed.Path = "/postgres"
	return dbName, parsed.String(), nil
}
