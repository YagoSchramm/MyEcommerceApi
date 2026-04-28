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
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	config := NewConfig()
	if config.cacheAddr == "" {
		config.cacheAddr = "localhost:6379"
	}
	if config.dbUrl == "" {
		config.dbUrl = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.dbUser, config.dbPassword, config.dbHost, config.dbPort, config.dbName, config.dbSSLMode)
	}

	if *migrateFlag {
		fmt.Println("Running database migrations...")
		if err := runMigrations(config.dbUrl); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
		fmt.Println("Migrations completed successfully")
		return nil
	}

	api := NewApi(config.dbUrl, config.cacheAddr, config.secret, config.addr)
	if err := api.Start(); err != nil {
		return fmt.Errorf("api startup failed: %w", err)
	}

	return nil
}
