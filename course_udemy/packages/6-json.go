package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonString := `{"name": "Igor", "age": 22}`
	var person Person
	json.Unmarshal([]byte(jsonString), &person)
	fmt.Printf("Nome: %s, Idade: %d \n", person.Name, person.Age)

	person.Name = "Adrieli"
	person.Age = 25

	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData))
}
