package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/angelhack2019/lib/models"
	"net/http"
)

var MockUser = &models.User{
	UUID:       1,
	Email:      "jdoe@uw.edu",
	Password:   "Unencrypted",
	SumRatings: 0,
	NumRatings: 0,
	Bio:        "I like food",
	School:     "UW Bothel",
	State:      "WA",
	Phone:      "2061234567",
	PicUrl:     "https://f4.bcbits.com/img/a0777435656_10.jpg",
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(MockUser)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Login")
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Auth")
}
