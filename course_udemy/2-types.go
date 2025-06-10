package main

import "fmt"

func main() {
	var idade int = 23
	var altura float64 = 1.83
	var maiorDeIdade bool = idade >= 18
	var nome string = "Igor"

	fmt.Println("Dados pessoais:")

	fmt.Println("Nome:")
	fmt.Println(nome)

	fmt.Println("Idade:")
	fmt.Println(idade)

	fmt.Println("Altura:")
	fmt.Println(altura)

	fmt.Println("Maior de idade?")
	fmt.Println(maiorDeIdade)

	//Mostrando o tipo da variável "maiorDeIdade"
	fmt.Println(fmt.Sprintf("%T", maiorDeIdade))

}
