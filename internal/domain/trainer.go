package domain

type Trainer struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Description string `json:"description"`
	Price       int    `json:"price" binding:"required"`
}
