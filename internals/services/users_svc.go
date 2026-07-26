package services

import (
	"backend/internals/libs"
	"backend/internals/models"
)

type UserServices struct{}

func (s *UserServices) GetAvailUsers(q *libs.QueryParams) []*models.UserResponse {
	return []*models.UserResponse{
		{
			Id:       325325,
			Username: "ranger",
			FullName: "Space Ranger",
		},
	}
}
