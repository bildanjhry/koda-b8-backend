package services

import (
	"backend/internals/libs"
	"backend/internals/models"
	"backend/internals/repository"
)

type UserServices struct {
	repo repository.UserRepository
}

func (s *UserServices) GetAvailUsers(q *libs.QueryParams) []*models.UserResponse {
	res := s.repo.GetAvailUsers(q)
	return res
}
