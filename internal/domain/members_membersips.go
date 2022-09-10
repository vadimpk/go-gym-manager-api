package domain

import "time"

type MembersMemberships struct {
	ID                   int       `json:"id" db:"id"`
	MemberID             int       `json:"member_id" db:"member_id"`
	MembershipID         int       `json:"membership_id" db:"membership_id"`
	MembershipExpiration time.Time `json:"membership_expiration" db:"membership_expiration"`
}
