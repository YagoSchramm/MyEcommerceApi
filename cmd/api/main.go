package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/handler"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	wd, _ := os.Getwd()

	uploadDir := filepath.Join(wd, "./images/")
	repo := repository.NewImageRepository(uploadDir)
	usecase := usecase.NewImageUsecase(repo, "http://localhost:8080")
	handler := handler.NewImageHandler(usecase)

	r.HandleFunc("/image/save", handler.Save)

	// servir arquivos
	fs := http.FileServer(http.Dir("./images/"))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))
	http.ListenAndServe(":8080", r)
}
