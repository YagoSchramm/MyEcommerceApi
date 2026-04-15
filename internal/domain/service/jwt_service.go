package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	secret string
}

func NewTokenService(secret string) *TokenService {
	return &TokenService{secret: secret}
}

type AccessClaims struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}
type RefreshClaims struct {
	jwt.RegisteredClaims
}

func (s *TokenService) GenerateTokens(userID string, roles []string) (string, string, error) {
	atClaims := AccessClaims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	accessToken, err := at.SignedString([]byte(s.secret))
	if err != nil {
		return "", "", err
	}

	rtClaims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	refreshToken, err := rt.SignedString([]byte(s.secret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
func (s *TokenService) ValidateAccessToken(tokenStr string) (*AccessClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccessClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func (s *TokenService) ValidateRefreshToken(tokenStr string) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
