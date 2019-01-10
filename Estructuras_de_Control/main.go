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

}
