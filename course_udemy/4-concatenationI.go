package main

import "fmt"

func main() {
	var nome string
	var idade int
	var altura float64
	var maiorDeIdade bool

	fmt.Println("Informe seu nome:")
	fmt.Scan(&nome)

	fmt.Println("Informe sua idade:")
	fmt.Scan(&idade)

	fmt.Println("Informe sua altura:")
	fmt.Scan(&altura)

	maiorDeIdade = idade >= 18
	fmt.Println("\nDados pessoais:\n")

	fmt.Printf("Nome: %s \n", nome)
	fmt.Printf("Idade: %d \n", idade)
	fmt.Printf("Altura: %.2f \n", altura)
	fmt.Printf("Maior de idade? %v \n", maiorDeIdade)
}
