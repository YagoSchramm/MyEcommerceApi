package service

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateID() *uuid.UUID {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil
	}
	return &id
}
func GenerateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	passwordHash := string(hashedPassword)
	return passwordHash, err
}
func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
func HasAnyRole(userRoles, allowedRoles []string) bool {
	roleSet := make(map[string]struct{}, len(userRoles))

	for _, r := range userRoles {
		roleSet[r] = struct{}{}
	}

	for _, allowed := range allowedRoles {
		if _, ok := roleSet[allowed]; ok {
			return true
		}
	}

	return false
}
func GetUserID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value("userID").(string)
	return id, ok
}

func GetRoles(ctx context.Context) ([]string, bool) {
	roles, ok := ctx.Value("roles").([]string)
	return roles, ok
}
func GetIp(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		ip := strings.TrimSpace(ips[0])
		if ip != "" {
			return ip
		}
	}

	xrip := r.Header.Get("X-Real-IP")
	if xrip != "" {
		return xrip
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return ip
}
