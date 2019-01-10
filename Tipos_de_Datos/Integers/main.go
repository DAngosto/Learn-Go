package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// Enteros con signo
	//var integer8 int8	// Todos los enteros de 8-bit con signo (-128 a 127)
	//var integer16 int16	// Todos los enteros de 16-bit con signo (-32768 a 32767)
	var integer32 int32 // Todos los enteros de 32-bit con signo (-2147483648 a 2147483647)
	var integer64 int64 // Todos los enteros de 64-bit con signo (-9223372036854775808 a 9223372036854775807)

	// Enteros sin signo
	//var uinteger8 int8	// Todos los enteros de 8-bit sin signo (0 a 255)
	//var uinteger16 int16	// Todos los enteros de 16-bit sin signo (0 a 65535)
	//var uinteger32 int32	// Todos los enteros de 32-bit sin signo (0 a 4294967295)
	//var uinteger64 int64	// Todos los enteros de 64-bit sin signo (0 a 18446744073709551615)

	//Alias
	//var integerByte byte 	// Alias para uint8
	var integerRune rune // Alias para int32

	// Tipos dependientes de la plataforma
	//var integerUint uint	// 32 o 64 bits
	var integerInt int // 32 o 64 bits
	//var integerUintptr uintptr	// Entero sin signo lo suficientemente grande para contener el valor de un puntero

	//***********************************************
	// Conversión entre tipos

	// Suma int32 y int64

	integer32 = 25
	integer64 = 85

	//fmt.Println(integer32 + integer64) // No se puede porque son tipos diferentes
	fmt.Println(integer32 + int32(integer64)) // Si deja porque le hacemos un casting a int32 al integer64

	//***********************************************
	// Suma int32 y rune

	integerRune = 46
	fmt.Println(integer32 + integerRune)

	//***********************************************
	// Suma int32 y int

	integerInt = 56
	//fmt.Println(integer32 + integerInt) // No se puede porque son tipos diferentes
	fmt.Println(integer32 + int32(integerInt))

	fmt.Println(unsafe.Sizeof(integerInt), unsafe.Sizeof(integer32)) // Devuelve el número de Bytes de cada tipo, en caso de 64bits 8 y 4, en 32bits 4 y 4

}
