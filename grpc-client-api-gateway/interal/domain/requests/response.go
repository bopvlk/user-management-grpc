package requests

import "time"

type DeleteAt struct {
	Time  time.Time `json:"deleted_at"`
	Valid bool      `json:"valid"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt
}
