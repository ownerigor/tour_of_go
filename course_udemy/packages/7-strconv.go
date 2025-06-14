package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Convertendo string para inteiro
	numeroStr := "123"
	numero, err := strconv.Atoi(numeroStr)
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	fmt.Println("Número: ", numero)

	//Convertendo inteiro para string
	numero2 := 456
	numeroStr2 := strconv.Itoa(numero2)
	fmt.Println(numeroStr2)

	//Convertendo string para float
	floatStr := "12.34"
	valorFloat, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	fmt.Println(valorFloat)
}
