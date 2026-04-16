package handler

import (
	"net/http"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase"
	"github.com/gorilla/mux"
)

type ImageHandler struct {
	usecase *usecase.ImageUsecase
}

func NewImageHandler(uc *usecase.ImageUsecase) *ImageHandler {
	return &ImageHandler{usecase: uc}
}
func (h *ImageHandler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/image/save", h.Save).Methods("POST")
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))
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
