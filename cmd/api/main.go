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
	config := NewConfig()
	// Database connection

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.dbUser, config.dbPassword, config.dbHost, config.dbPort, config.dbName, config.dbSSLMode)

	if *migrateFlag {
		fmt.Println("Running database migrations...")
		if err := runMigrations(connStr); err != nil {
			fmt.Printf("Migration failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Migrations completed successfully")
		return
	}

	api := NewApi(connStr, config.secret, config.addr, config.cacheAddr)
	if err := api.Start(); err != nil {
		panic(err)
	}
}
