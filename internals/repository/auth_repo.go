package repository

import "backend/internals/models"

type AuthRepository struct{}

func (r *AuthRepository) Register(form *models.RegisterRequest) (*models.RegisterResponse, error) {
	return &models.RegisterResponse{
		Id: 32523523,
	}, nil
}

func (r *AuthRepository) Login(form *models.LoginRequest) (*models.LoginResponse, error) {
	return &models.LoginResponse{
		Id:    3235235325,
		Token: "sdgsdgwegweg",
	}, nil
}
