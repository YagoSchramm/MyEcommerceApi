package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	usc *usecase.UserUsecase
}

func NewUserHandler(usc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usc: usc}
}

func (h *UserHandler) MountPublicHandlers(r *mux.Router) {
	r.HandleFunc("/users", h.CreateUser).Methods("POST")
}

func (h *UserHandler) MountProtectedHandlers(r *mux.Router) {
	r.HandleFunc("/users/{id}", h.GetUserById).Methods("GET")
	r.HandleFunc("/users", h.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createDTO dto.CreateUserDTO
	err := json.NewDecoder(r.Body).Decode(&createDTO)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.usc.CreateUser(r.Context(), &createDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := h.usc.GetUserById(r.Context(), &dto.GetUserByIdDTO{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.usc.GetAllUsers(r.Context(), &dto.GetAllUsersDTO{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	var updateDTO dto.UpdateUserDTO
	err := json.NewDecoder(r.Body).Decode(&updateDTO)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	updateDTO.ID = idStr

	err = h.usc.UpdateUser(r.Context(), &updateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	err := h.usc.DeleteUser(r.Context(), &dto.DeleteUserDTO{ID: idStr})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
