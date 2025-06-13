package main

import "fmt"

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco string
}

func main() {
	pessoa := Pessoa{
		Nome:     "Igor Queirantes",
		Idade:    22,
		Endereco: "Av. das Araras, 123",
	}
	fmt.Println("Informações da pessoa")
	fmt.Printf("Nome: %s \n", pessoa.Nome)
	fmt.Printf("Idade: %d \n", pessoa.Idade)
	fmt.Printf("Endereço: %s \n", pessoa.Endereco)
}
