package main

import "fmt"

// Em Go pode-se retornar múltiplos valores em uma função
func swap(x, y string) (string, string) {
	return y, x
}

func multipleResults() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
