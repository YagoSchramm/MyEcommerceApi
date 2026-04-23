package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/service"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	productUsc *usecase.ProductUsecase
	userUsc    *usecase.UserUsecase
}

func NewProductHandler(productUsc *usecase.ProductUsecase, userUsc *usecase.UserUsecase) *ProductHandler {
	return &ProductHandler{productUsc: productUsc, userUsc: userUsc}
}

func (h *ProductHandler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/products", h.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", h.GetProductById).Methods("GET")
	r.HandleFunc("/products", h.GetAllProducts).Methods("GET")
	r.HandleFunc("/products/{id}", h.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", h.DeleteProduct).Methods("DELETE")
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var createDTO dto.CreateProductDTO
	err := json.NewDecoder(r.Body).Decode(&createDTO)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	userIDStr, ok := service.GetUserID(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	// Get user name
	user, err := h.userUsc.GetUserById(r.Context(), &dto.GetUserByIdDTO{ID: userID})
	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	createDTO.UserID = userID
	createDTO.UserName = user.Name

	id, err := h.productUsc.CreateProduct(r.Context(), &createDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id.String()})
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	product, err := h.productUsc.GetProductById(r.Context(), &dto.GetProductByIdDTO{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productUsc.GetAllProducts(r.Context(), &dto.GetAllProductsDTO{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	var updateDTO dto.UpdateProductDTO
	err := json.NewDecoder(r.Body).Decode(&updateDTO)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	updateDTO.ID, _ = uuid.Parse(idStr)

	err = h.productUsc.UpdateProduct(r.Context(), &updateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := uuid.Parse(idStr)
	err := h.productUsc.DeleteProduct(r.Context(), &dto.DeleteProductDTO{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
