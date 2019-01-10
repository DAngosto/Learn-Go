package main

import "fmt"

var razaGokuGlobal = "Saiyajin Global" //Variable Global

func main() {
	// Nombrar variables
	// Acepta letras, números y _
	// Se recomienda hacer lowerCamelCase en el caso de unir varias palabras
	// Distingue entre mayúsculas y minúsculas
	numero := 25
	_nombre := "Dani"
	numero2 := 54
	nombreUsuario := "Admin"
	Numero := 1
	fmt.Println(numero, numero2, _nombre, nombreUsuario, Numero)
	// Crear varias variables al mismo tiempo mediante el uso de var
	var (
		dios               = "Goku"
		enemigo1, enemigo2 = "Babidi", "Celula"
	)
	var numero3 = 66
	fmt.Println(dios, enemigo1, enemigo2, numero3)

	// Scope
	var razaGoku = "Saiyajin" // La variable es solo local, solo puede ser usada por la función main
	fmt.Println("La raza de Goku es: " + razaGoku)

	imprimir()

}

func imprimir() {
	//fmt.Println("La raza de Goku es: " + razaGoku)  // razaGoku es local, no nos lo permite
	fmt.Println("La raza de Goku es: " + razaGokuGlobal)

}
