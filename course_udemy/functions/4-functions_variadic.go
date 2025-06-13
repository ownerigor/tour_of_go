package main

import "fmt"

// 1 - Função para soma de números com variádicos
func sum(nums ...int) {
	sumTotal := 0
	for _, n := range nums {
		sumTotal += n
	}
	fmt.Printf("Soma é: %d \n", sumTotal)
}

// 2 - Função para apresentação de cursos com variádicos
func presentationData(data map[string]string) {
	for key, value := range data {
		fmt.Printf("%s = %s \n", key, value)
	}
}

func main() {
	sum(7)
	sum(7, 3)
	sum(7, 3, 5)
	sum(7, 3, 5, 10)
	sum(7, 3, 5, 10, 2)

	presentationData(map[string]string{
		"name":     "Python",
		"category": "Backend",
		"level":    "Iniciante",
	})

	presentationData(map[string]string{
		"name":     "Visão computacional com Python",
		"category": "IA",
		"level":    "Avançado",
	})

	presentationData(map[string]string{
		"name":     "Dashboards com Dash",
		"category": "Data Science",
		"level":    "Intermediário",
	})
}
