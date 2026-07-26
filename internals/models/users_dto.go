package models

import "time"

type UserResponse struct {
	Id          int64     `json:"id"`
	Username    string    `json:"usersname"`
	FullName    string    `json:"fullname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	BirthDate   time.Time `json:"brith_date"`
	Gender      int       `json:"gender"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
