package repository

import (
	"backend/internals/libs"
	"backend/internals/models"
)

type UserRepository struct{}

func (r *UserRepository) GetAvailUsers(q *libs.QueryParams) []*models.UserResponse {
	return []*models.UserResponse{
		{
			Id:       325325,
			Username: "ranger",
			FullName: "Space Ranger",
		},
	}
}
