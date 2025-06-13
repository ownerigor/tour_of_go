package main

import "fmt"

func main() {
	//Criando um map com nome do aluno como chave e a nota como valor
	estudantes := map[string]float64{
		"Ana":     7.5,
		"Carlos":  4.3,
		"Beatriz": 8.9,
		"João":    6.2,
	}
	fmt.Println("Classificação dos alunos:")
	for nome, nota := range estudantes {
		status := "Reprovado"
		if nota >= 6.0 {
			status = "Aprovado"
		}
		fmt.Printf("Aluno: %s | Nota: %.2f | Status: %s \n", nome, nota, status)
	}
}
