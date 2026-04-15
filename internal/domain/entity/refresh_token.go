package entity

import "time"

type RefreshToken struct {
	UserID string
	Token  string
	Expiry time.Time
}
