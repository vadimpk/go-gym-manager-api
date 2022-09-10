package domain

import "time"

type MembersVisits struct {
	ID        int       `json:"id" db:"id"`
	MemberID  int       `json:"member_id" db:"member_id"`
	ManagerID int       `json:"manager_id" db:"manager_id"`
	ArrivedAt time.Time `json:"arrived_at" db:"arrived_at"`
	LeftAt    time.Time `json:"left_at" db:"left_at"`
}

type TrainersVisits struct {
	ID        int       `json:"id" db:"id"`
	ManagerID int       `json:"manager_id" db:"manager_id"`
	TrainerID int       `json:"trainer_id" db:"trainer_id"`
	ArrivedAt time.Time `json:"arrived_at" db:"arrived_at"`
	LeftAt    time.Time `json:"left_at" db:"left_at"`
}
