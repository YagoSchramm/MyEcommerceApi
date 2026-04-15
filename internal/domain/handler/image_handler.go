package handler

import (
	"net/http"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase"
)

type ImageHandler struct {
	usecase *usecase.ImageUsecase
}

func NewImageHandler(uc *usecase.ImageUsecase) *ImageHandler {
	return &ImageHandler{usecase: uc}
}

func (h *ImageHandler) Save(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "erro ao processar", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "arquivo inválido", http.StatusBadRequest)
		return
	}
	defer file.Close()

	img, err := h.usecase.Save(file, handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(img.Url))
}
