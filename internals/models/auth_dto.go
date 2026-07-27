package models

type LoginRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterRequest struct {
	FullName string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	Id    int64 `json:"id"`
	Token string
}

type RegisterResponse struct {
	Id int64 `json:"id"`
}
