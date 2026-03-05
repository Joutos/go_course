package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // Connection drive
)

// Connect - Connect to the database.
//
// This function connects to the database using the environment
// variables DB_USER, DB_PASS, DB_HOST and DB_DATABASE.
//
// It returns a pointer to the sql.DB object and an error.
// If the connection was successful, the error is nil.
// If there was an error connecting to the database, the
// error is non-nil and the sql.DB pointer is nil.
func Connect() (*sql.DB, error) {

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")

	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, database)

	db, erro := sql.Open("mysql", connString)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
