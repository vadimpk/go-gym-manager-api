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

type MembersMembership struct {
	ID           int       `json:"id" db:"id"`
	MemberID     int       `json:"member_id" db:"member_id"`
	MembershipID int       `json:"membership_id" db:"membership_id"`
	ExpiresAt    time.Time `json:"membership_expiration" db:"membership_expiration"`
}

type MembersMembershipResponse struct {
	Membership Membership `json:"membership"`
	ExpiresAt  time.Time  `json:"membership_expiration"`
}
