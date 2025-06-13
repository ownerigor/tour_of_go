package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func welcome() {
	fmt.Println("Bem vindo ao sistema de filmes!")
}

func createMovie() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Digite o nome do filme:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Digite o ano do filme:")
	var yearRelease int
	fmt.Scanln(&yearRelease)

	fmt.Println("Digite o preço do filme:")
	var moviePrice float64
	fmt.Scanln(&moviePrice)

	fmt.Printf("%s (%d) - R$ %.2f\n", name, yearRelease, moviePrice)
}

func calculateAverage() float64 {
	var numRatings int
	fmt.Println("Digite quantas avaliações deseja fazer para o filme:")
	fmt.Scan(&numRatings)

	var total float64
	for i := 0; i < numRatings; i++ {
		var note float64
		fmt.Println("Digite a nota para o filme:")
		fmt.Scan(&note)
		total += note
	}

	var average float64

	if numRatings > 0 {
		average = total / float64(numRatings)
	} else {
		average = 0
	}
	return average
}

func main() {
	welcome()
	fmt.Println("Utilizando função")
	createMovie()
	media := calculateAverage()
	fmt.Printf("A média das avaliações é: %.2f \n", media)
}
