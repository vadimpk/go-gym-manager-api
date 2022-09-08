package domain

import "time"

type MembersVisits struct {
	ID        int       `json:"id"`
	MemberID  int       `json:"member_id" binding:"required"`
	ManagerID int       `json:"manager_id"`
	TrainerID int       `json:"trainer_id"`
	ArrivedAt time.Time `json:"arrived_at"`
	LeftAt    time.Time `json:"left_at"`
}

type TrainersVisits struct {
	ID        int       `json:"id"`
	ManagerID int       `json:"manager_id"`
	TrainerID int       `json:"trainer_id" binding:"required"`
	ArrivedAt time.Time `json:"arrived_at"`
	LeftAt    time.Time `json:"left_at"`
}
