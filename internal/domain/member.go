package domain

import "time"

type Member struct {
	ID                   int       `json:"id"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	PhoneNumber          string    `json:"phone_number"`
	MembershipID         int       `json:"membership_id"`
	JoinedAt             time.Time `json:"joined_at"`
	MembershipExpiration time.Time `json:"membership_expiration"`
}
