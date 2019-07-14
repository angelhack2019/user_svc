package main

import (
	"github.com/angelhack2019/user_svc/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	initRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initRoutes() {
	router := mux.NewRouter().StrictSlash(true)
	router.Host("localhost").Schemes("http")

	router.HandleFunc("/login", controllers.HandleLogin).Methods("POST")
	router.HandleFunc("/user/{uuid}", controllers.HandleGetUser).Methods("GET")
	router.HandleFunc("/user/{uuid}", controllers.HandleEditUser).Methods("PUT")
	router.HandleFunc("/user", controllers.HandleNewUser).Methods("POST")

	http.Handle("/", router)
}
