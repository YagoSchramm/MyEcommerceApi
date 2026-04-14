package rules

import (
	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity/derr"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service/dto"
)

func ValidateCreateProduct(dto dto.CreateProductDTO) error {
	if !isValidUserID(dto.UserID.String()) {
		return derr.InvalidUserIDErr
	}
	if !isValidName(dto.Name) {
		return derr.InvalidNameErr
	}
	if !isValidPrice(dto.Value) {
		return derr.InvalidPriceErr
	}
	if !isValidStock(int(dto.Stock)) {
		return derr.InvalidStockErr
	}
	if !isValidDescription(dto.Description) {
		return derr.InvalidDescriptionErr
	}
	return nil
}
func ValidateUpdateProduct(dto dto.UpdateProductDTO) error {
	return nil
}
func isValidPrice(price float32) bool {
	return price > 0
}
func isValidStock(stock int) bool {
	return stock >= 0
}
func isValidUserID(userID string) bool {
	return userID != ""
}
func isValidName(name string) bool {
	return name != "" && len(name) >= 2
}
func isValidDescription(description string) bool {
	return description != "" && len(description) >= 10 && len(description) <= 1000
}
