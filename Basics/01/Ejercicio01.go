package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	fmt.Println("Primer programa en Go!!")
	fmt.Println("Ingeresa un numero: ")
	fmt.Scanln(&n1)
	fmt.Println("Ingeresa otro numero: ")
	fmt.Scanln(&n2)
	fmt.Println("La suma de los numeros es: ", sum(n1, n2))
}

func sum(a int, b int) int {
	return a + b
}
