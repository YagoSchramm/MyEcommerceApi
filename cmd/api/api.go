package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/handler"
	domainMiddleware "github.com/YagoSchramm/myecommerce-api/internal/domain/middleware"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase"
	"github.com/YagoSchramm/myecommerce-api/internal/foundation"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
	infraMiddleware "github.com/YagoSchramm/myecommerce-api/internal/infrastructure/middleware"
	"github.com/gorilla/mux"
)

type Api struct {
	connStr   string
	secret    string
	addr      string
	cacheAddr string
}

func NewApi(connStr, cacheAddr, secret, addr string) *Api {
	return &Api{
		connStr:   connStr,
		secret:    secret,
		addr:      addr,
		cacheAddr: cacheAddr,
	}
}

func (api *Api) Start() error {
	r := mux.NewRouter()
	// Initialize logger
	logger, err := foundation.NewLogger()
	if err != nil {
		return err
	}
	defer logger.Sync()

	// Apply logging middleware to all routes
	r.Use(infraMiddleware.LoggingMiddleware(logger))

	db, err := foundation.NewPostgresDB(api.connStr)
	if err != nil {
		return err
	}
	defer db.Close()
	rdb := foundation.NewClient(api.cacheAddr)
	tokenService := service.NewTokenService(api.secret)
	rateRepo := repository.NewRedisLimiter(rdb, 10, time.Second)
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	ratingRepo := repository.NewRatingRepository(db)
	purchaseRepo := repository.NewPurchaseRepository(db)
	imageRepo := repository.NewImageRepository("./uploads")

	userUsecase := usecase.NewUserUsecase(userRepo, tokenService)
	productUsecase := usecase.NewProductUsecase(productRepo)
	ratingUsecase := usecase.NewRatingUsecase(ratingRepo)
	purchaseUsecase := usecase.NewPurchaseUsecase(purchaseRepo)
	imageUsecase := usecase.NewImageUsecase(imageRepo, "http://localhost"+api.addr)

	authHandler := handler.NewAuthHandler(userUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	productHandler := handler.NewProductHandler(productUsecase, userUsecase)
	ratingHandler := handler.NewRatingHandler(ratingUsecase, userUsecase)
	purchaseHandler := handler.NewPurchaseHandler(purchaseUsecase)
	imageHandler := handler.NewImageHandler(imageUsecase)
	rateLimiting := domainMiddleware.NewRateLimitMiddleware(rateRepo)
	authHandler.MountHandlers(r)
	userHandler.MountPublicHandlers(r)
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(rateLimiting.Handler)
	protected.Use(domainMiddleware.AuthMiddleware(tokenService))
	imageHandler.MountHandlers(protected)
	userHandler.MountProtectedHandlers(protected)
	productHandler.MountHandlers(protected)
	ratingHandler.MountHandlers(protected)
	purchaseHandler.MountHandlers(protected)

	wd, _ := os.Getwd()
	uploadDir := filepath.Join(wd, "./uploads")
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadDir))))
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if port == "" {
		port = "8080"
	}
	return http.ListenAndServe(port, r)
}
