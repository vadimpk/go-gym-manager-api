package domain

type Membership struct {
	ID          int    `json:"id" db:"id"`
	ShortName   string `json:"short_name" db:"short_name"`
	Description string `json:"description" id:"description"`
	Price       int    `json:"price" db:"price"`
	Duration    string `json:"duration" db:"duration"`
}

type MembershipCreateInput struct {
	ShortName   string `json:"short_name" binding:"required"`
	Description string `json:"description"`
	Price       int    `json:"price" binding:"required"`
	Duration    string `json:"duration" binding:"required"`
}

type MembershipUpdateInput struct {
	ShortName   string `json:"short_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Duration    string `json:"duration"`
}
