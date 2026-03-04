package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloWorld)

	http.HandleFunc("/users", showUsers)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func showUsers(w http.ResponseWriter, r *http.Request) {
	users := []string{"João", "Sofia", "Lucas"}
	jsonUsers, erro := json.Marshal(users)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUsers)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
