package main

import (
	"fmt"
	"strings"
)

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
	var sb strings.Builder

	fmt.Println("\nDados pessoais:\n")

	//Outra forma de dar input em informações com strings.Builder
	sb.WriteString(fmt.Sprintf("Nome: %s \n", nome))
	sb.WriteString(fmt.Sprintf("Idade: %d \n", idade))
	sb.WriteString(fmt.Sprintf("Altura: %.2f \n", altura))
	sb.WriteString(fmt.Sprintf("Maior de idade? %v \n", maiorDeIdade))

	fmt.Println(sb.String())
}
