package rules

import (
	"path/filepath"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity/derr"
)

func ValidateImageFile(filename string) error {
	ext := filepath.Ext(filename)
	allowed := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	if !allowed[ext] {
		return derr.InvalidImageErr
	}
	return nil
}
