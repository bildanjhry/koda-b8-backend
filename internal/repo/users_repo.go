package repo

import (
	"backend/internal/model"
	"errors"
)

type UserRepo struct {
	data *[]model.Users
}

func (r *UserRepo) NewUserRepo(data *[]model.Users) *UserRepo {
	return &UserRepo{
		data: data,
	}
}

func (r *UserRepo) Create(data *model.Users) *model.Users {
	id := len(*r.data) + 1
	*r.data = append(*r.data, model.Users{
		Id:       id,
		Email:    data.Email,
		Password: data.Password,
	})
	return &model.Users{
		Id:    data.Id,
		Email: data.Email,
	}
}

func Create(data *model.UserForm) (*model.Users, error) {
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
