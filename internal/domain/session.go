package domain

import "time"

type Session struct {
	ID           int       `json:"id" db:"id"`
	RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at" db:"expires_at"`
	ManagerID    int       `json:"manager_id" db:"manager_id"`
}

type RefreshInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
