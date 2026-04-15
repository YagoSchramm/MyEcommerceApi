package usecase

import (
	"errors"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service"
)

type TokenService interface {
	ValidateRefreshToken(tokenStr string) (*service.RefreshClaims, error)
	GenerateTokens(userID string, roles []string) (string, string, error)
}

type RefreshRepository interface {
	Save(userID, token string) error
	Exists(userID, token string) bool
	Delete(userID, token string) error
}

type UserRepository interface {
	GetByID(userID string) (*entity.User, error)
}

type RefreshUseCase struct {
	tokenService TokenService
	refreshRepo  RefreshRepository
	userRepo     UserRepository
}

func NewRefreshUseCase(tokenService TokenService, refreshRepo RefreshRepository, userRepo UserRepository) *RefreshUseCase {
	return &RefreshUseCase{
		tokenService: tokenService,
		refreshRepo:  refreshRepo,
		userRepo:     userRepo,
	}
}

// Execute validates the refresh token, checks database, rotates the token and returns new access token and refresh token
func (uc *RefreshUseCase) Execute(rt string) (newAT, newRT string, err error) {
	// Validate JWT
	claims, err := uc.tokenService.ValidateRefreshToken(rt)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	userID := claims.Subject

	// Verify token exists in database
	if !uc.refreshRepo.Exists(userID, rt) {
		return "", "", errors.New("refresh token not found or revoked")
	}

	// Get user to retrieve roles
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	// Convert roles to strings
	roles := make([]string, len(user.Roles))
	for i, r := range user.Roles {
		roles[i] = string(r)
	}

	// Generate new tokens
	newAT, newRT, err = uc.tokenService.GenerateTokens(userID, roles)
	if err != nil {
		return "", "", err
	}

	// Delete old refresh token from database
	err = uc.refreshRepo.Delete(userID, rt)
	if err != nil {
		return "", "", err
	}

	// Save new refresh token to database
	err = uc.refreshRepo.Save(userID, newRT)
	if err != nil {
		return "", "", err
	}

	return newAT, newRT, nil
}
