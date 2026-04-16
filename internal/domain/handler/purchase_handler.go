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

type PurchaseHandler struct {
	usc *usecase.PurchaseUsecase
}

func NewPurchaseHandler(usc *usecase.PurchaseUsecase) *PurchaseHandler {
	return &PurchaseHandler{usc: usc}
}

func (h *PurchaseHandler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/purchases", h.CreatePurchase).Methods("POST")
	r.HandleFunc("/purchases/{id}", h.GetPurchaseById).Methods("GET")
	r.HandleFunc("/purchases/user/{userId}", h.GetAllPurchaseByUserId).Methods("GET")
	r.HandleFunc("/purchases", h.GetAllPurchases).Methods("GET")
}

func (h *PurchaseHandler) CreatePurchase(w http.ResponseWriter, r *http.Request) {
	var createDTO dto.CreatePurchaseDTO
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

	createDTO.UserID = userID

	id, err := h.usc.CreatePurchase(r.Context(), &createDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id.String()})
}

func (h *PurchaseHandler) GetPurchaseById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid purchase id", http.StatusBadRequest)
		return
	}

	purchase, err := h.usc.GetPurchaseById(r.Context(), &dto.GetPurchaseByIdDTO{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(purchase)
}

func (h *PurchaseHandler) GetAllPurchaseByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdStr := vars["userId"]
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	purchases, err := h.usc.GetAllPurchaseByUserId(r.Context(), &dto.GetAllPurchaseByUserIdDTO{UserID: userId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(purchases)
}

func (h *PurchaseHandler) GetAllPurchases(w http.ResponseWriter, r *http.Request) {
	purchases, err := h.usc.GetAllPurchases(r.Context(), &dto.GetAllPurchasesDTO{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(purchases)
}
