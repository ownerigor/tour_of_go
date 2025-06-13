package main

import (
	"fmt"
	"strings"
)

func main() {
	//Criando um map para contar as palavras
	text := "go é uma linguagem de programação e go é muito rápido, ela é parecida com C++"
	words := strings.Fields(text)
	fmt.Println(words)

	wordCount := make(map[string]int)

	//Contagem da frequência de palavras
	for _, word := range words {
		wordCount[word]++
	}

	//Imprimir as frequências
	fmt.Println("Contagem de palavras")
	for word, count := range wordCount {
		fmt.Printf("Palavra: %s | Frequência: %d \n", word, count)
	}
}
