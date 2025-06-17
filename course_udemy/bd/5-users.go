package main

import (
	"database/sql"
	"fmt"
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

func insertUser(db *sql.DB, user *User) error {
	stmt, err := db.Prepare("INSERT INTO users (id, name, email, age) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func updateUser(db *sql.DB, user *User) error {
	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Email, user.Age, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func deleteUser(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM users where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
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

	//Inserção de usuário
	user := newUser("John Doe", "john.doe@example.com", 30)
	err = insertUser(db, user)
	if err != nil {
		log.Fatal(err)
	}

	//Alterando usuário
	user.Name = "John Smith"
	user.Age = 31
	err = updateUser(db, user)
	if err != nil {
		log.Fatal(err)
	}

	//Listagem de usuários
	users, err := selectUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Printf("Usuário %s possui email %s e possui %d anos\n", u.Name, u.Email, u.Age)
	}

	//Exclusão de dados
	err = deleteUser(db, user.ID)
	if err != nil {
		log.Fatal(err)
	}
}
