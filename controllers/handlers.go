package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/angelhack2019/lib/mocks"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(mocks.MockUsers[0])
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Login")
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Auth")
}
