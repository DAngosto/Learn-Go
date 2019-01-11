package main

import (
	"fmt"
)

//Declarar Funciones 1

func imprimirNombre(nombre string) {
	fmt.Println("Fuera de Main")
	fmt.Println("El nombre es: ", nombre)
}

//Declarar Funciones 2
func suma(n1 int, n2 int) int {
	return n1 + n2
}

//Declarar Funciones 3
func resta(n1, n2 int) (r int) {
	r = n1 - n2
	return
}

//Funciones que retornan multiples valores

func multiplicar(numero int) (n1, n2, n3 int) {
	n1 = numero * 10
	n2 = numero * 20
	n3 = numero * 30
	return
}

func multiplicar2(numero int) (int, int, int) {
	n1 := numero * 10
	n2 := numero * 20
	n3 := numero * 30
	return n1, n2, n3
}

func retorno() (string, string) {

	return "Hola", "Mundo"
}

func main() {
	//Funciones

	//Llamar una Funcion
	imprimirNombre("Jose")
	fmt.Println("Dentro de main")

	fmt.Println(suma(25, 66))
	fmt.Println(resta(88, 66))
	//
	//Firma de una funcion
	fmt.Printf("Funcion suma: %T\n", suma)
	fmt.Printf("Funcion resta: %T\n", resta)

	//package math
	//func Sin(x float64) float64 // implemented in assembly language

	fmt.Println(multiplicar(25))
	c1, c2, c3 := multiplicar(65)
	fmt.Println(c1, c2, c3)
	fmt.Println(multiplicar2(96))
	_, d2, _ := multiplicar(66)
	fmt.Println(d2)
	fmt.Println(retorno())

}
