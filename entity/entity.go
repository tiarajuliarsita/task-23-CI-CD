package entity

import "time"

type Core struct {
	ID        uint
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
