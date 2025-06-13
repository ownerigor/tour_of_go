package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&num1)

	fmt.Println("Digite outro número:")
	fmt.Scan(&num2)

	//Operadores aritméticos
	soma := num1 + num2
	sub := num1 - num2
	mult := num1 * num2
	div := num1 / num2
	mod := num1 % num2

	fmt.Printf("Soma de %d e %d é %d \n", num1, num2, soma)
	fmt.Printf("Subtração de %d e %d é %d \n", num1, num2, sub)
	fmt.Printf("Multiplicação de %d e %d é %d \n", num1, num2, mult)
	fmt.Printf("Divisão de %d e %d é %d \n", num1, num2, div)
	fmt.Printf("Resto da divisão de %d e %d é %d \n", num1, num2, mod)

	//Atribuição
	num1 += 1 //num1 = num1 + 1
	num1 -= 1 //num1 = num1 - 1
	num1 *= 1 //num1 = num1 * 1
	num1 /= 1 //num1 = num1 / 1

	fmt.Println("Valor final de num1 após alterações de atribuição: ", num1)

	//Comparação
	bigger := num1 > num2
	smaller := num1 < num2
	equal := num1 == num2
	different := num1 != num2
	biggerEqual := num1 >= num2
	smallerEqual := num1 <= num2

	fmt.Printf("O número %d é maior que o número %d? %v \n", num1, num2, bigger)
	fmt.Printf("O número %d é menor que o número %d? %v \n", num1, num2, smaller)
	fmt.Printf("O número %d é igual ao número %d? %v \n", num1, num2, equal)
	fmt.Printf("O número %d é diferente ao número %d? %v \n", num1, num2, different)
	fmt.Printf("O número %d é maior ou igual que o número %d? %v \n", num1, num2, biggerEqual)
	fmt.Printf("O número %d é menor ou igual que o número %d? %v \n", num1, num2, smallerEqual)
}
