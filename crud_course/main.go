package main

import (
	"log"
	"net/http"
	user "simple_crud/controller"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", user.GetUsers).Methods("GET")
	router.HandleFunc("/usuarios", user.CreateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
