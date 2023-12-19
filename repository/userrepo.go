package repository

import (
	"restapi/model"
)

func GetById(id string) (*model.User, error) {
	var user *model.User
	user.Id = id
	user.Password = "password" + id
	user.Name = "Android No:" + id
	user.Email = id + "@example.com"
	return user, nil
}
