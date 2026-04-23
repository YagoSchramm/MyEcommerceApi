package repository

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type ImageRepository struct {
	uploadDir string
}

func NewImageRepository(uploadDir string) *ImageRepository {
	return &ImageRepository{uploadDir: uploadDir}
}

func (r *ImageRepository) Save(file multipart.File, filename string) (string, error) {

	imagesDir := filepath.Join(r.uploadDir)

	if err := os.MkdirAll(imagesDir, 0755); err != nil {
		return "", err
	}
	path := filepath.Join(imagesDir, filename)
	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return "/uploads/" + filename, nil
}
