package controllers

import (
	"github.com/angelhack2019/lib/models"
	"github.com/angelhack2019/lib/utility"
	"github.com/angelhack2019/user_svc/interfaces"
	"github.com/gorilla/mux"
	"net/http"
)

type Credentials struct {
	Email    string `json:email`
	Password string `json:password`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	utility.ParseRequestJSON(r, creds)

	if creds.Email == "" || creds.Password == "" {
		utility.RespondWithError(w, http.StatusBadRequest, "Missing email or password field")
		return
	}

	uuid, ok := interfaces.LoginUser(creds.Email, creds.Password)

	if ok {
		utility.RespondWithJSON(w, http.StatusOK, uuid)
		return
	} else {
		utility.RespondWithError(w, http.StatusUnauthorized, "Invalid login")
	}
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	user, ok := interfaces.GetUser(uuid)

	if !ok {
		utility.RespondWithError(w, http.StatusOK, "Failed to get user with uuid: "+uuid)
		return
	}

	utility.RespondWithJSON(w, http.StatusOK, user)
}

func HandleEditUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utility.ParseRequestJSON(r, user)

	ok := interfaces.EditUser(user)

	if !ok {
		utility.RespondWithError(w, http.StatusOK, "Failed to update user with uuid: "+uuid)
		return
	}

	utility.Respond(w, http.StatusOK, "Success")
}

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utility.ParseRequestJSON(r, user)

	uuid, ok := interfaces.CreateUser(user)

	if !ok {
		utility.RespondWithError(w, http.StatusOK, "Error creating new user "+user.Email)
		return
	}

	utility.RespondWithJSON(w, http.StatusOK, uuid)
}
