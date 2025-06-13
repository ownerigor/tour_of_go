package main

import (
	"fmt"
	"strings"
)

func main() {
	movieName := "Top Gun"
	movieName2 := "top Gun"

	//A linguagem Go é case sensitive
	fmt.Println(movieName == movieName2)

	movieDescription :=
		`Top Gun Maveric é um filme de aviação e
aventura, muito consagrado na indústria`

	//Convertendo texto
	fmt.Println(strings.ToUpper(movieDescription))
	fmt.Println(strings.ToLower(movieDescription))
	fmt.Println(strings.Title(movieDescription))

	//Encontrar posição de um caracterere
	fmt.Println(strings.Index(movieName, "p"))

	//Contando o número de ocorrências de um caracterere
	fmt.Println(strings.Count(movieDescription, "a"))
	fmt.Println(strings.Count(movieDescription, "e"))

	//Substituindo um elemento por outro
	fmt.Println(strings.ReplaceAll(movieDescription, "filme", "série"))

	//Divindo a string
	fmt.Println(strings.Split(movieDescription, ","))
}
