package main

import (
	"fmt"
)

func main() {
	//Range

	nombres := []string{
		"Alejandro",
		"Maria",
		"Pedro",
		"Carlos",
	}

	for index, nombre1 := range nombres {
		// Index = 0
		// nombre nombres[index]
		fmt.Printf("El nombre %q esta en el index %d \n", nombre1, index)
	}

	for _, nombre2 := range nombres {
		fmt.Println(nombre2)
	}

	for index1 := range nombres {
		fmt.Println(index1)
	}

	cadena := "123 responde otra vez"

	for index2, letra := range cadena {
		fmt.Printf("La letra %q esta en el index %d \n", letra, index2)
	}

}
