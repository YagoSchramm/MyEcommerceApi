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

type RatingHandler struct {
	ratingUsc *usecase.RatingUsecase
	userUsc   *usecase.UserUsecase
}

func NewRatingHandler(ratingUsc *usecase.RatingUsecase, userUsc *usecase.UserUsecase) *RatingHandler {
	return &RatingHandler{ratingUsc: ratingUsc, userUsc: userUsc}
}

func (h *RatingHandler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/ratings", h.CreateRating).Methods("POST")
	r.HandleFunc("/ratings/{id}", h.GetRatingById).Methods("GET")
	r.HandleFunc("/ratings/user/{userId}", h.GetRatingByUserId).Methods("GET")
	r.HandleFunc("/ratings/product/{productId}", h.GetAllByProductId).Methods("GET")
	r.HandleFunc("/ratings/{id}", h.UpdateRating).Methods("PUT")
	r.HandleFunc("/ratings/{id}", h.DeleteRating).Methods("DELETE")
}

func (h *RatingHandler) CreateRating(w http.ResponseWriter, r *http.Request) {
	var createDTO dto.CreateRatingDTO
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

	err = h.ratingUsc.CreateRating(r.Context(), &createDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *RatingHandler) GetRatingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid rating id", http.StatusBadRequest)
		return
	}

	rating, err := h.ratingUsc.GetRatingById(r.Context(), &dto.GetRatingByIdDTO{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rating)
}

func (h *RatingHandler) GetRatingByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdStr := vars["userId"]
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	ratings, err := h.ratingUsc.GetRatingByUserId(r.Context(), &dto.GetRatingByUserIdDTO{UserID: userId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ratings)
}

func (h *RatingHandler) GetAllByProductId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productIdStr := vars["productId"]
	productId, err := uuid.Parse(productIdStr)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	ratings, err := h.ratingUsc.GetAllByProductId(r.Context(), &dto.GetAllRatingByProductIdDTO{ProductID: productId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ratings)
}

func (h *RatingHandler) UpdateRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid rating id", http.StatusBadRequest)
		return
	}

	var updateDTO dto.UpdateRatingDTO
	err = json.NewDecoder(r.Body).Decode(&updateDTO)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	updateDTO.ID = id

	err = h.ratingUsc.UpdateRating(r.Context(), &updateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *RatingHandler) DeleteRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid rating id", http.StatusBadRequest)
		return
	}

	err = h.ratingUsc.DeletRating(r.Context(), &dto.DeleteRatingDTO{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
