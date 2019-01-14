package main

import (
	"fmt"
	"time"
)

func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//Generamos los numeros
	go func() {
		for x := 0; x < 5; x++ {
			numero <- x
		}
		close(numero)
	}()

	//Lo elevamos al cuadrado
	go func() {
		for x := range numero { //Utilizamos un for range para que una vez se acabe de mandar datos por el canal, este se cierre
			cuadrado <- x * x
		}
		close(cuadrado)
	}()

	//Imprimimos en main
	for x := range cuadrado {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}

}
