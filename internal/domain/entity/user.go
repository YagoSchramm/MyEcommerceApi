package entity

import "github.com/google/uuid"

type Role string

const (
	RoleBuyer  Role = "buyer"
	RoleSeller Role = "seller"
	RoleAdmin  Role = "admin"
)

type User struct {
	ID       uuid.UUID `json:"user_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Roles    []Role    `json:"roles"`
}

// Domain Rules
func (u *User) HasRole(role Role) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}
