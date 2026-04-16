package main

import (
	"flag"
	"fmt"
	"os"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	// Database connection
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "user")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "myecommerce")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	if *migrateFlag {
		fmt.Println("Running database migrations...")
		if err := runMigrations(connStr); err != nil {
			fmt.Printf("Migration failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Migrations completed successfully")
		return
	}

	secret := getEnv("JWT_SECRET", "secret")
	addr := getEnv("API_ADDR", ":8080")

	api := NewApi(connStr, secret, addr)
	if err := api.Start(); err != nil {
		panic(err)
	}
}
