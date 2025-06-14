package main

import (
	"fmt"
	"math"
)

// 1 - Acessar o número PI

func acessarPI() {
	fmt.Printf("PI arredondado (2 casas decimais): %.2f \n", math.Pi)
}

// 2 - Acessar o número Euler
func acessarEuler() {
	fmt.Printf("Euler arredondado (2 casas decimais): %.2f \n", math.E)
}

// 3 - Arredondando números para cima e para baixo
func arredondar() {
	num := 10.4
	fmt.Println("Arredondando para cima: ", math.Ceil(num))
	fmt.Println("Arredondando para baixo: ", math.Floor(num))
}

func main() {
	acessarPI()
	acessarEuler()
	arredondar()
	fmt.Println("Potência de 5 elevado 5: ", math.Pow(5, 5))
	fmt.Println("Raiz quadrada de 169: ", math.Sqrt(169))
	fmt.Println("Logaritmo de 10: ", math.Log(10))
}
