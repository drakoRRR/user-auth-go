package models

import "time"

type RegisterUserPayload struct {
	Name     string `json:"first_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"first_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"first_name"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
