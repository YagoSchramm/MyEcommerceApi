package main

import "github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"

type Store struct {
	*repository.UserRepository
	*repository.ProductRepository
	*repository.RatingRepository
	*repository.PurchaseRepository
}

func NewStore(userRepo *repository.UserRepository, productRepo *repository.ProductRepository, ratingRepo *repository.RatingRepository, purchaseRepo *repository.PurchaseRepository) *Store {
	return &Store{
		UserRepository:     userRepo,
		ProductRepository:  productRepo,
		RatingRepository:   ratingRepo,
		PurchaseRepository: purchaseRepo,
	}
}
