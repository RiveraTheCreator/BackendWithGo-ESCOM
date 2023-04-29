package main

import (
	"fmt"
)

func main() {
	var n1, n2, opcion int
	fmt.Println("Ingeresa un numero: ")
	fmt.Scanln(&n1)
	fmt.Println("Ingeresa otro numero: ")
	fmt.Scanln(&n2)
	opcion = imprimirMenu()
	calculadora(n1, n2, opcion)

}

func imprimirMenu() int {
	var opcion int
	fmt.Println("-----Mi programa en Go!! -----")
	fmt.Println("-----Calculadora -----")
	fmt.Print("\nIngeresa la operacion: \n1 suma, \n2 resta \n3 multiplicacion \n4 division\n5 potencia")
	fmt.Scanln(&opcion)
	return opcion
}

func calculadora(n1 int, n2 int, opcion int) int {
	//TODO: implementar
	switch opcion {
	case 1:
		fmt.Println("La suma de los numeros es: ", sum(n1, n2))
	case 2:
		fmt.Println("La resta de los numeros es: ", resta(n1, n2))
	case 3:
		fmt.Println("La multiplicacion de los numeros es: ", multiplicacion(n1, n2))
	case 4:
		fmt.Println("La division de los numeros es: ", division(n1, n2))
	case 5:
		fmt.Println("La potencia de los numeros es: ", potencia(n1, n2))
	default:
		fmt.Println("Opcion no valida")
	}

	return 0
}

func sum(a int, b int) int {
	return a + b
}

func resta(a int, b int) int {
	return a - b
}

func multiplicacion(a int, b int) int {
	return a * b
}

func division(a int, b int) int {
	return a / b
}

func potencia(a int, b int) int {
	return a ^ b
}
