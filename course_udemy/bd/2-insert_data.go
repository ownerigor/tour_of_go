package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO users (name) VALUES (?)", "Igor Queirantes")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Novo usuário inserido com sucesso!")
}
