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

	router.HandleFunc("/", controllers.HandleHome).Methods("GET")
	router.HandleFunc("/login", controllers.HandleLogin).Methods("POST")

	http.Handle("/", router)
}
