package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")

	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, database)

	db, erro := sql.Open("mysql", connString)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	lines, _ := db.Query("SELECT * from usuarios")
	defer lines.Close()

	for lines.Next() {
		var id int
		var nome string
		var email string

		err := lines.Scan(&id, &nome, &email)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, nome, email)
	}
}
