package main

import "fmt"

func main() {

	// Estructura For

	// Imprimir los numeros del 1 al 5
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++

	}

	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("555")
		break
	}

	//Estructuras de Control: if

	// if 5 < 6 {
	// 	fmt.Println("5 es menor que 6")
	// }

	a := 6

	if a < 6 {
		fmt.Println("a es menor que 6")
	} else if a > 6 {
		fmt.Println("a es mayor que 6")
	} else {
		fmt.Println("a es igual 6")
	}

	fmt.Println("Fin del programa")

}
