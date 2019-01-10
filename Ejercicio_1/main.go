package main

import "fmt"

func main() {

	//Contador de Números Impares

	encabezado := `
	****************************
	Contador de Numeros Impares
	****************************
	`
	//Imprimimos el encabezado
	fmt.Println(encabezado)
	fmt.Println("Introduce el Primer número")
	var numero1 int
	fmt.Scanln(&numero1)
	fmt.Println("Introduce el Segundo número")
	var numero2 int
	fmt.Scanln(&numero2)
	var contador int
	for i := numero1; i <= numero2; i++ {
		//Evaluamos si el numero es impar
		if i%2 != 0 {
			//Si el numero es impar incrementamos el valor de la variable contador en 1
			contador++
			//Imprimimos el numero impar
			fmt.Printf("%d es un numero impar\n", i)
		}
	}

	encabezado = `
	*********************
	Resultado del Conteo
	*********************
	`
	//Imprimimos el encabezado
	fmt.Println(encabezado)
	fmt.Printf("Entre el %d y el %d hay %d numeros impares.", numero1, numero2, contador)

}
