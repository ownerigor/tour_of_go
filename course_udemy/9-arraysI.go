package main

import "fmt"

func main() {
	//Declarando um array com 5 elementos do tipo int
	var numeros [5]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	fmt.Println("Array de números: ", numeros)
}
