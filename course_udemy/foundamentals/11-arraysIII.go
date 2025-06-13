package main

import "fmt"

func main() {
	var notas = [5]float64{7.5, 8.3, 9.0, 5.7, 6.4}

	soma := notas[0] + notas[1] + notas[2] + notas[3] + notas[4]
	media := soma / float64(len(notas))

	fmt.Println("A soma das notas é: ", soma)
	fmt.Printf("A média dessas notas é: %.2f \n", media)
}
