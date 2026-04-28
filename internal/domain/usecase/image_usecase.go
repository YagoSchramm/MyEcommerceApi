package usecase

import (
	"mime/multipart"
	"path/filepath"

	"github.com/YagoSchramm/myecommerce-api/internal/domain"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/rules"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/google/uuid"
)

type ImageUsecase struct {
	repo    domain.ImageRepository
	baseUrl string
}

func NewImageUsecase(repo domain.ImageRepository, baseUrl string) *ImageUsecase {
	return &ImageUsecase{repo: repo, baseUrl: baseUrl}
}
func (uc *ImageUsecase) Save(file multipart.File, filename string) (*dto.ImageResponseDTO, error) {
	err := rules.ValidateImageFile(filename)
	if err != nil {
		return nil, err
	}
	ext := filepath.Ext(filename)
	newFileName := uuid.New().String() + ext
	path, err := uc.repo.Save(file, newFileName)
	if err != nil {
		return nil, err
	}
	return &dto.ImageResponseDTO{
		Url:  uc.baseUrl + path,
		Path: path,
	}, nil
}
