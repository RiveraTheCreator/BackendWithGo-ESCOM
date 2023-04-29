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
	if sum(n1, n2) > 10 {
		fmt.Println("La suma es mayor a 10 ", "\U0001F600")
	} else {
		fmt.Println("La suma es menor a 10", "\U0001F61E")
	}

}

func sum(a int, b int) int {
	return a + b
}
