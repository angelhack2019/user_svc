package interfaces

import (
	"github.com/angelhack2019/lib/mocks"
	"github.com/angelhack2019/lib/models"
)

func LoginUser(email string, password string) (uuid string, ok bool) {

	// Todo: Logic for checking password against database and fetching ID

	return mocks.MockUsers[0].UUID, true
}

func GetUser(uuid string) (user models.User, ok bool) {

	// Todo: Query for User row from UUID

	return mocks.MockUsers[0], true
}

func CreateUser(user *models.User) (uuid string, ok bool) {

	// Todo: Query for insert User row from User model

	return "Created new user: " + user.Email, true
}

func EditUser(user *models.User) (ok bool) {

	// Todo: Query for update User where UUID

	uuid := user.UUID

	return true
}