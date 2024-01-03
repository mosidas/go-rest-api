package repository

import (
	"restapi/model"
)

type IUserRepository interface {
	GetById(id string) (*model.User, error)
}

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (ur userRepository) GetById(id string) (*model.User, error) {
	user := &model.User{
		Credential: model.Credential{
			Id:       id,
			Password: "password" + id,
		},
		Name:  "Android No:" + id,
		Email: id + "@example.com",
	}
	println(user)
	return user, nil
}
