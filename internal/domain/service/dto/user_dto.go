package dto

import (
	"time"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/util"
	"github.com/google/uuid"
)

type CreateUserDTO struct {
	Name     string        `json:"name" validate:"required"`
	Email    string        `json:"email" validate:"required,email"`
	Password string        `json:"password" validate:"required,min=6"`
	Roles    []entity.Role `json:"roles"`
}

type UserResponseDTO struct {
	ID    uuid.UUID     `json:"id"`
	Name  string        `json:"name"`
	Email string        `json:"email"`
	Roles []entity.Role `json:"roles"`
}
type UpdateUserDTO struct {
	ID    string         `json:"id"`
	Name  *string        `json:"name,omitempty"`
	Roles *[]entity.Role `json:"roles,omitempty"`
}

func ToUserEntity(dto CreateUserDTO) *entity.User {
	now := time.Now()
	return &entity.User{
		ID:        *util.GenerateID(),
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  dto.Password,
		Roles:     dto.Roles,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func ToUserResponseDTO(user *entity.User) UserResponseDTO {
	return UserResponseDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Roles: user.Roles,
	}
}
