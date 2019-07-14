package models

import (
	"github.com/angelhack2019/lib/mocks"
	"github.com/angelhack2019/lib/models"
)

func LoginUser(email string, password string) (uuid string, ok bool) {
	return mocks.MockUsers[0].UUID, true
}

func GetUser(uuid string) (user models.User, ok bool) {
	return mocks.MockUsers[0], true
}

func CreateUser(user interface{}) (uuid string, ok bool) {
	return mocks.MockUsers[0].UUID, true
}

func EditUser(uuid string) (ok bool) {
	return true
}