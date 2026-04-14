package rules

import (
	"regexp"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity/derr"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service/dto"
)

func ValidateCreateUser(user *dto.CreateUserDTO) error {
	if user.Name == "" || len(user.Name) < 2 {
		return derr.InvalidNameErr
	}
	if user.Email == "" || len(user.Email) < 5 || !isValidEmail(user.Email) {
		return derr.InvalidEmailErr
	}
	if user.Password == "" {
		return derr.EmptyPasswordErr
	}
	if !isValidPassword(user.Password) || len(user.Password) < 6 {
		return derr.InvalidPasswordErr
	}
	for _, role := range user.Roles {
		if !isValidRole(string(role)) {
			return derr.InvalidRoleErr
		}
	}
	return nil
}
func ValidateUpdateUser(updateIt *dto.UpdateUserDTO) error {
	if updateIt.Name != nil && (*updateIt.Name) == "" {
		return derr.InvalidNameErr
	}
	if updateIt.Name != nil && len(*updateIt.Name) < 2 {
		return derr.InvalidNameErr
	}
	if updateIt.Roles != nil {
		for _, role := range *updateIt.Roles {
			if !isValidRole(string(role)) {
				return derr.InvalidRoleErr
			}
		}
	}
	return nil
}
func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}
func isValidPassword(password string) bool {
	const passwordRegex = `^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{6,}$`
	return regexp.MustCompile(passwordRegex).MatchString(password)
}
func isValidRole(role string) bool {
	return role == "admin" || role == "customer" || role == "seller"
}
