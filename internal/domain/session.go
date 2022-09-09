package domain

import "time"

type Session struct {
	ID           int       `json:"id"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}
