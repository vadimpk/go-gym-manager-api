package domain

type Manager struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Email       string `json:"email"`
}

type SignInInput struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
