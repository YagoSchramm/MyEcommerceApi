package rules

import (
	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity/derr"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
)

func ValidateCreatePurchase(dto dto.CreatePurchaseDTO) error {
	if !isValidUserID(dto.UserID.String()) {
		return derr.InvalidUserIDErr
	}
	if !isValidProductID(dto.ProductID.String()) {
		return derr.InvalidProductIDErr
	}
	if !isValidQuantity(dto.Quantity) {
		return derr.InvalidQuantityErr
	}
	return nil
}
func isValidProductID(productID string) bool {
	return productID != ""
}
func isValidQuantity(quantity int) bool {
	return quantity > 0
}
