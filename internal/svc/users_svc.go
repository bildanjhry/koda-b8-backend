package svc

import (
	"backend/internal/model"
	"errors"
)

type UserService struct {
	data *[]model.Users
}

func (r *UserService) NewUserService(data *[]model.Users) *UserService {
	return &UserService{
		data: data,
	}
}

func (r *UserService) Create(data *model.UserForm) (*model.Users, error) {
	id := 1
	if data.Email == "admin@mail.com" && data.Password == "1234" {

		return &model.Users{
			Id:    id,
			Email: data.Email,
		}, nil

	} else {
		return &model.Users{}, errors.New("Invalid")
	}
}
