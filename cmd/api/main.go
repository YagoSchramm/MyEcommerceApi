package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/handler"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/middleware"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase"
	"github.com/YagoSchramm/myecommerce-api/internal/foundation"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
	"github.com/gorilla/mux"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	r := mux.NewRouter()

	// Database
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "user")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "myecommerce")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbName, dbSSLMode)
	db, err := foundation.NewPostgresDB(connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Services
	tokenService := service.NewTokenService("secret")

	// Repositories
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	ratingRepo := repository.NewRatingRepository(db)
	purchaseRepo := repository.NewPurchaseRepository(db)
	imageRepo := repository.NewImageRepository("./uploads")

	// Usecases
	userUsecase := usecase.NewUserUsecase(userRepo, tokenService)
	productUsecase := usecase.NewProductUsecase(productRepo)
	ratingUsecase := usecase.NewRatingUsecase(ratingRepo)
	purchaseUsecase := usecase.NewPurchaseUsecase(purchaseRepo)
	imageUsecase := usecase.NewImageUsecase(imageRepo, "http://localhost:8080")

	// Handlers
	authHandler := handler.NewAuthHandler(userUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	productHandler := handler.NewProductHandler(productUsecase, userUsecase)
	ratingHandler := handler.NewRatingHandler(ratingUsecase, userUsecase)
	purchaseHandler := handler.NewPurchaseHandler(purchaseUsecase)
	imageHandler := handler.NewImageHandler(imageUsecase)

	// Mount handlers
	authHandler.MountHandlers(r)
	userHandler.MountHandlers(r)
	imageHandler.MountHandlers(r)

	// Protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware(tokenService))

	productHandler.MountHandlers(protected)
	ratingHandler.MountHandlers(protected)
	purchaseHandler.MountHandlers(protected)

	// Static files
	wd, _ := os.Getwd()
	uploadDir := filepath.Join(wd, "./uploads")
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadDir))))

	http.ListenAndServe(":8080", r)
}
