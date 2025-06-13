package main

import (
	"fmt"
)

func main() {
	var numero int

	fmt.Println("Informe um número:")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Println("O número é positivo")
	} else if numero < 0 {
		fmt.Println("O número é negativo")
	} else {
		fmt.Println("O número é 0 ou inválido")
	}

	if numero%2 == 0 {
		fmt.Println("O número é par")
	} else {
		fmt.Println("O número é ímpar")
	}
}
