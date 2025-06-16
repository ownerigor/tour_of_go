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

	//Atualizando dados
	result, err := db.Exec("UPDATE users SET name = ? WHERE id = ?", "Matteo", 3)
	if err != nil {
		log.Fatal(err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Update %d row(s) \n", affectedRows)

	//Exclusão de dados
	result, err = db.Exec("DELETE FROM users WHERE id = ?", 2)
	if err != nil {
		log.Fatal(err)
	}
	affectedRows, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Delete %d row(s) \n", affectedRows)
}
