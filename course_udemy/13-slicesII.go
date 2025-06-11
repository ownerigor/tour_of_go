package main

import "fmt"

func main() {
	frutas := []string{"Maçã", "Banana", "Laranja", "Uva", "Morango"}

	//Criar um subslice, pegando do índice de 1 a 3
	subslice := frutas[1:4]
	fmt.Println("Subslice de frutas: ", subslice)

	subslice[0] = "Manga"
	fmt.Println("Subslice de frutas após a modificação: ", subslice)
	fmt.Println("Slice de frutas após a modificação: ", frutas)
}
