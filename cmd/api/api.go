package main

import (
	"net/http"
	"os"
	"path/filepath"

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
	connStr string
	secret  string
	addr    string
}

func NewApi(connStr, secret, addr string) *Api {
	return &Api{
		connStr: connStr,
		secret:  secret,
		addr:    addr,
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

	tokenService := service.NewTokenService(api.secret)

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

	authHandler.MountHandlers(r)
	userHandler.MountPublicHandlers(r)

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(domainMiddleware.AuthMiddleware(tokenService))
	imageHandler.MountHandlers(protected)
	userHandler.MountProtectedHandlers(protected)
	productHandler.MountHandlers(protected)
	ratingHandler.MountHandlers(protected)
	purchaseHandler.MountHandlers(protected)

	wd, _ := os.Getwd()
	uploadDir := filepath.Join(wd, "./uploads")
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadDir))))

	return http.ListenAndServe(api.addr, r)
}
