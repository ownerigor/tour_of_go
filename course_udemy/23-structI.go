package main

import "fmt"

// Definindo a struct Carro
type Carro struct {
	Modelo string
	Ano    int
	Cor    string
}

func main() {
	//Criar a instância da struct Carro
	carro1 := Carro{
		Modelo: "Fusca",
		Ano:    1965,
		Cor:    "Azul",
	}

	fmt.Println("Informações do carro")
	fmt.Printf("Modelo: %s \n", carro1.Modelo)
	fmt.Printf("Ano: %d \n", carro1.Ano)
	fmt.Printf("Cor: %s \n", carro1.Cor)
}
