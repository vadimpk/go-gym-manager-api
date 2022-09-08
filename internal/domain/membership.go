package domain

import "time"

type Membership struct {
	ID          int           `json:"id"`
	ShortName   string        `json:"short_name" binding:"required"`
	Description string        `json:"description"`
	Price       int           `json:"price" binding:"required"`
	Duration    time.Duration `json:"duration" binding:"required"`
}
