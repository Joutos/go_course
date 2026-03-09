package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"simple_crud/db"
	"strconv"

	"github.com/gorilla/mux"
)

type Usuario struct {
	UserId    int    `json:"userId"`
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
}

// GetUsers - Retorna todos os usuários do banco de dados.
//
// Esta função retorna todos os usuários do banco de dados na forma de um JSON.
// Se houver um erro ao conectar ao banco de dados, a função retorna um erro HTTP com status 500.
// Se houver um erro ao executar o statement, a função retorna um erro HTTP com status 500.
// Se houver um erro ao obter os usuários, a função retorna um erro HTTP com status 500.
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

// GetUser - Retorna um usuário do banco de dados.
//
// Esta função recebe um ID por parâmetro e retorna o usuário correspondente
// em formato JSON.
// Se houver um erro ao conectar ao banco de dados, a função retorna um erro HTTP com status 500.
// Se houver um erro ao executar o statement, a função retorna um erro HTTP com status 500.
// Se houver um erro ao obter o usuário, a função retorna um erro HTTP com status 500.
// Se o usuário não for encontrado, a função retorna um erro HTTP com status 404.
func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	db, err := db.Connect()
	if err != nil {
		http.Error(w, "Falha ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u Usuario

	err = db.QueryRow(
		"SELECT userId, userName, userEmail FROM usuarios WHERE userId = ?",
		id,
	).Scan(&u.UserId, &u.UserName, &u.UserEmail)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// CreateUser - Cria um novo usuário no banco de dados.
//
// Esta função recebe um JSON com informações do usuário a ser criado e
// retorna o ID do usuário criado.
// Se houver um erro ao conectar ao banco de dados, a função retorna um erro HTTP com status 500.
// Se houver um erro ao executar o statement, a função retorna um erro HTTP com status 500.
// Se houver um erro ao obter o ID do usuário, a função retorna um erro HTTP com status 500.
// Se o JSON recebido não tiver as informações necessárias, a função retorna um erro HTTP com status 400.
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

// PutUser - Atualiza um usuário no banco de dados.
//
// Esta função recebe um ID de usuário e um JSON com informações do usuário a ser atualizado.
// Se houver um erro ao conectar ao banco de dados, a função retorna um erro HTTP com status 500.
// Se houver um erro ao executar o statement, a função retorna um erro HTTP com status 500.
// Se houver um erro ao obter o ID do usuário, a função retorna um erro HTTP com status 500.
// Se o JSON recebido não tiver as informações necessárias, a função retorna um erro HTTP com status 400.
func PutUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), 400)
		w.Write([]byte("ID inválido"))
		return
	}
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	requestBody, _ := io.ReadAll(r.Body)

	var body Usuario
	var u Usuario

	json.Unmarshal(requestBody, &body)

	if body.UserName == "" && body.UserEmail == "" {
		http.Error(w, "Alguma coisa precisa ser alterada", 400)
		return
	}

	err = db.QueryRow(
		"SELECT userId, userName, userEmail FROM usuarios WHERE userId = ?",
		id,
	).Scan(&u.UserId, &u.UserName, &u.UserEmail)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if u.UserName != body.UserName {
		u.UserName = body.UserName
	}
	if u.UserEmail != body.UserEmail {
		u.UserEmail = body.UserEmail
	}

	stmt, err := db.Prepare("UPDATE usuarios SET userName = ?, userEmail = ? WHERE userId = ?")
	if err != nil {
		http.Error(w, err.Error(), 500)
		w.Write([]byte("Erro ao executar o statement"))
		return
	}
	_, err = stmt.Exec(u.UserName, u.UserEmail, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		w.Write([]byte("Erro ao executar o statement"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Usuário alterado com sucesso! ID: %d", id)))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), 400)
		w.Write([]byte("ID inválido"))
		return
	}
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM usuarios WHERE userId = ?")
	if err != nil {
		http.Error(w, err.Error(), 500)
		w.Write([]byte("Erro ao executar o statement"))
		return
	}
	_, err = stmt.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		w.Write([]byte("Erro ao executar o statement"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Usuário deletado com sucesso! ID: %d", id)))
}
