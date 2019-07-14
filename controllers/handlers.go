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
	utility.ParseRequestJSON(w, r, creds)

	if creds.Email == "" || creds.Password == "" {
		utility.RespondWithError(w, http.StatusBadRequest, "Missing email or password field")
		return
	}

	uuid, ok := interfaces.LoginUser(creds.Email, creds.Password)

	if ok {
		utility.RespondWithJSON(w, http.StatusOK, uuid)
		return
	} else {
		msg := uuid
		utility.RespondWithError(w, http.StatusUnauthorized, msg)
	}
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	user, errorMsg := interfaces.GetUser(uuid)

	if errorMsg != "" {
		utility.RespondWithError(w, http.StatusOK, "Failed to get user with uuid: "+uuid)
		return
	}

	utility.RespondWithJSON(w, http.StatusOK, user)
}

func HandleEditUser(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	user := &models.User{}
	utility.ParseRequestJSON(w, r, user)

	user.UUID = uuid

	errorMsg := interfaces.EditUser(user)

	if errorMsg != "" {
		utility.RespondWithError(w, http.StatusOK, errorMsg)
		return
	}

	utility.Respond(w, http.StatusOK, "Success")
}

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utility.ParseRequestJSON(w, r, user)

	uuid, errorMsg := interfaces.CreateUser(user)

	if errorMsg != "" {
		utility.RespondWithError(w, http.StatusOK, errorMsg)
		return
	}

	utility.RespondWithJSON(w, http.StatusOK, uuid)
}
