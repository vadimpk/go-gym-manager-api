package domain

import "time"

type Member struct {
	ID          int       `json:"id" db:"id"`
	FirstName   string    `json:"first_name" db:"first_name" binding:"required"`
	LastName    string    `json:"last_name" db:"last_name"  binding:"required"`
	PhoneNumber string    `json:"phone_number" db:"phone_number" binding:"required"`
	JoinedAt    time.Time `json:"joined_at" db:"joined_at"`
}

type MemberCreateInput struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type MemberUpdateInput struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}
