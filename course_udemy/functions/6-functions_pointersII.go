package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
}

func atualizarCliente(c *Cliente, novoNome string, novaIdade int) {
	c.Nome = novoNome
	c.Idade = novaIdade
}

func main() {
	cliente := Cliente{Nome: "Igor", Idade: 22}
	fmt.Println("Antes da alteração: ", cliente)
	atualizarCliente(&cliente, "Carlos", 28)
	fmt.Println("Após a alteração: ", cliente)
}
