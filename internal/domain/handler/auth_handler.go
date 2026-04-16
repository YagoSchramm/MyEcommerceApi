package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/gorilla/mux"
)

type AuthHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewAuthHandler(userUsecase *usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{userUsecase: userUsecase}
}

func (h *AuthHandler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/auth/login", h.Login).Methods("POST")
	r.HandleFunc("/auth/logout", h.Logout).Methods("POST")
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginDTO dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&loginDTO)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.userUsecase.Login(r.Context(), &loginDTO)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// For now, just return success. In a real app, you might blacklist the token.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("logged out"))
}
