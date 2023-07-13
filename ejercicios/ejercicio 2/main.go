package main

import (
	"fmt"
)

func main() {
	/*
		Debemos implementar un programa en go para ir ingresando valores por
		consola hasta que se agrega cero y finaliza el programa.

		Todos los valores ingresados por consola se deben agregar a un
		array e imprimirlo por pantalla al finalizar.
	*/
	var array []int
	var value int = 1

	for value != 0 {
		fmt.Println("Ingresa un valor, 0 para salir")
		fmt.Scanln(&value)
		if value != 0 {
			array = append(array, value)
		}
	}

	fmt.Println("El valor del array es:", array)
}
