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
aventura muito consagrado na indústria`

	line := "-"
	fmt.Println(strings.Repeat(line, 50))
	fmt.Println(movieDescription)

	//Verifica se uma palavra existe dentro de uma string
	fmt.Println(strings.Contains(movieDescription, "top"))
	fmt.Println(strings.Contains(movieDescription, "filme"))
}
