package main

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID    string
	Name  string
	Email string
	Age   int
}

func newUser(name, email string, age int) *User {
	return &User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
		Age:   age,
	}
}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT,
		age INTEGER
	);
	`
	_, err := db.Exec(query)
	return err
}

func main() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Criar a tabela
	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}
}
