package main

import "fmt"

func main() {
	nomes := [4]string{"Ana", "Carlos", "Beatriz", "João"}
	fmt.Println("Estrutura de repetição for")

	for id, nome := range nomes {
		//fmt.Println("Id:", id)
		//fmt.Println("Nome:", nome)
		fmt.Println(id+1, "->", nome)
		if len(nome) > 5 {
			fmt.Println(nome, " tem mais de 5 caracteres")
		}
	}
}
