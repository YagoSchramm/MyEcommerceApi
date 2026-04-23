package rules

import (
	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity/derr"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
)

func CreateRating(dto dto.CreateRatingDTO) error {
	if !isValidUserName(dto.UserName) {
		return derr.InvalidNameErr
	}
	if !isValidDescription(dto.Description) {
		return derr.InvalidDescriptionErr
	}
	if !isValidRating(dto.Rating) {
		return derr.InvalidRatingErr
	}
	return nil
}
func UpdateRating(dto dto.UpdateRatingDTO) error {
	if !isValidRating(dto.Rating) {
		return derr.InvalidRatingErr
	}
	return nil
}
func isValidUserName(name string) bool {
	return name != "" && len(name) >= 2
}
func isValidRating(rating float32) bool {
	return rating >= 0 && rating <= 5
}
