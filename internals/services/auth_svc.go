package services

import (
	"backend/internals/models"
	"backend/internals/repository"
)

type AuthServices struct {
	repo repository.AuthRepository
}

func (s *AuthServices) Register(form *models.RegisterRequest) (*models.RegisterResponse, error) {
	res, err := s.repo.Register(form)
	return res, err
}

func (s *AuthServices) Login(form *models.LoginRequest) (*models.LoginResponse, error) {
	res, err := s.repo.Login(form)
	return res, err
}
