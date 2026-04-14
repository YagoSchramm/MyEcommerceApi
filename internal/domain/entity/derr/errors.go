package derr

import "errors"

// User related errors
var (
	InvalidNameErr     = errors.New("invalid name: name cannot be empty and must be at least 2 characters long")
	InvalidEmailErr    = errors.New("invalid email: email cannot be empty")
	EmptyPasswordErr   = errors.New("invalid password: password cannot be empty")
	InvalidPasswordErr = errors.New("invalid password: password must be at least 6 characters long and contain at least one letter and one number")
	InvalidRoleErr     = errors.New("invalid role: role must be either 'admin' or 'customer' or 'seller'")
)

// Product related errors
var (
	InvalidPriceErr       = errors.New("invalid price: price must be greater than 0")
	InvalidStockErr       = errors.New("invalid stock: stock cannot be negative")
	InvalidUserIDErr      = errors.New("invalid user ID: user ID cannot be empty")
	InvalidProductNameErr = errors.New("invalid product name: name cannot be empty and must be at least 2 characters long")
	InvalidDescriptionErr = errors.New("invalid description: description cannot be empty and must be between 10 and 1000 characters long")
)

// Purchase related errors
var (
	InvalidProductIDErr = errors.New("invalid product ID: product ID cannot be empty")
	InvalidQuantityErr  = errors.New("invalid quantity: quantity must be greater than 0")
)

// Rating related errors
var (
	InvalidRatingErr = errors.New("invalid Rating: rating cannot be empty or be at least 0 and the maximum of 5")
)
