package derr

import "errors"

// User related errors
var (
	EncondingUserErr   = errors.New("error on decoding the user response to json")
	DecondingUserErr   = errors.New("error on encoding the user from json")
	InvalidNameErr     = errors.New("invalid name: name cannot be empty and must be at least 2 characters long")
	InvalidEmailErr    = errors.New("invalid email: email cannot be empty")
	EmptyPasswordErr   = errors.New("invalid password: password cannot be empty")
	InvalidPasswordErr = errors.New("invalid password: password must be at least 6 characters long and contain at least one letter and one number")
	InvalidRoleErr     = errors.New("invalid role: role must be either 'admin' or 'customer' or 'seller'")
)

// Product related errors
var (
	EncondingProductErr   = errors.New("error on decoding the product response to json")
	DecondingProductErr   = errors.New("error on encoding the product from json")
	InvalidPriceErr       = errors.New("invalid price: price must be greater than 0")
	InvalidStockErr       = errors.New("invalid stock: stock cannot be negative")
	InvalidUserIDErr      = errors.New("invalid user ID: user ID cannot be empty")
	InvalidProductNameErr = errors.New("invalid product name: name cannot be empty and must be at least 2 characters long")
	InvalidDescriptionErr = errors.New("invalid description: description cannot be empty and must be between 10 and 1000 characters long")
)

// Purchase related errors
var (
	EncondingPurchaseErr = errors.New("error on decoding the purchase response to json")
	DecondingPurchaseErr = errors.New("error on encoding the purchase from json")
	InvalidProductIDErr  = errors.New("invalid product ID: product ID cannot be empty")
	InvalidQuantityErr   = errors.New("invalid quantity: quantity must be greater than 0")
)

// Rating related errors
var (
	EncondingRatingErr = errors.New("error on decoding the rating response to json")
	DecondingRatingErr = errors.New("error on encoding the rating from json")
	InvalidRatingErr   = errors.New("invalid Rating: rating cannot be empty or be at least 0 and the maximum of 5")
)
var (
	InvalidImageErr = errors.New("invalid image type or path")
)
