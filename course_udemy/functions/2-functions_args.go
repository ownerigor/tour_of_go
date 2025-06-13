package main

import "fmt"

func fullName(firstName, lastName string) {
	fmt.Printf("Nome completo: %s %s \n", firstName, lastName)
}

func sumNumbers(a, b int) int {
	return a + b
}

func address(country string) {
	if country == "" {
		country = "Brasil"
	}
	fmt.Printf("Eu moro no %s \n", country)
}

func main() {
	fmt.Println("Utilizando função com parâmetros")
	fullName("Igor", "Queirantes")
	fullName("Adrieli", "Brito")
	fmt.Printf("Soma: %d \n", sumNumbers(200, 50))
	address("")
	address("Canadá")
}
