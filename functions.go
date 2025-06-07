package main

import "fmt"

var number1, number2 int

func add(x int, y int) int {
	return x + y
}

func sumNumber() {
	fmt.Print("Digite um número: ")
	fmt.Scan(&number1)

	fmt.Print("Digite outro número: ")
	fmt.Scan(&number2)

	fmt.Println("A soma desses dois número são: ", add(number1, number2))
}
