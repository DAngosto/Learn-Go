package main

import "fmt"

func main() {
	var numero int // Hace que numero sea del tipo int y además le asigna el valor por defecto que ese tipo tiene, en este caso un 0.
	fmt.Println(numero)
	multiplicando := 25
	fmt.Println(numero * multiplicando)
	numero = 40
	fmt.Println(numero)

	nombre := "Alejandro" // Para crear una variable sin indicar el tipo
	nombre = "Pedro"

	nombre, numero = "Rafael", 88
	nombre2 := "Ramon"
	fmt.Println(nombre, nombre2)
	nombre, nombre2 = nombre2, nombre
	fmt.Println(nombre, nombre2)
	fmt.Println(numero)
	nombre3, nombre := "Maria", "Alejandro"
	fmt.Println(nombre3, nombre)
}
