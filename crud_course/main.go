package main

import (
	"fmt"
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
	router.HandleFunc("/usuarios/{id}", user.GetUser).Methods("GET")
	router.HandleFunc("/usuarios", user.CreateUser).Methods("POST")
	router.HandleFunc("/usuarios/{id}", user.PutUser).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", user.DeleteUser).Methods("DELETE")
	fmt.Printf("Server is running on port %s\n", "3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
