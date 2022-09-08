package domain

type Program struct {
	ID          int    `json:"id"`
	ShortName   string `json:"short_name" binding:"required"`
	Description string `json:"description"`
	Price       int    `json:"price" binding:"required"`
	TrainerID   int    `json:"trainer_id" binding:"required"`
}

type ProgramMembers struct {
	ID        int `json:"id"`
	MemberID  int `json:"member_id" binding:"required"`
	ProgramID int `json:"program_id" binding:"required"`
}
