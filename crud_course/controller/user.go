package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"simple_crud/db"
)

type Usuario struct {
	UserId    int    `json:"userId"`
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT userId, userName, userEmail FROM usuarios")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var u Usuario

		err := rows.Scan(&u.UserId, &u.UserName, &u.UserEmail)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		usuarios = append(usuarios, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := io.ReadAll(r.Body)

	var u Usuario
	json.Unmarshal(requestBody, &u)

	if u.UserName == "" || u.UserEmail == "" {
		http.Error(w, "userName and userEmail are required", 400)
		return
	}

	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), 500)
		w.Write([]byte("Falha ao conectar com o banco de dados"))
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO usuarios (userName, userEmail) VALUES (?, ?)")
	res, err := stmt.Exec(u.UserName, u.UserEmail)
	if err != nil {
		http.Error(w, err.Error(), 500)
		w.Write([]byte("Erro ao executar o statement"))
		return
	}
	insertedId, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), 500)
		w.Write([]byte("Erro ao obter o ID do usuário"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário criado com sucesso! ID: %d", insertedId)))
}
