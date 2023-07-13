package main

import "fmt"

func main() {
	/*
		Dentro de nuestro código ya tenemos definido un array, debemos reccorrerlo e incrementar todos sus valores en 20.
		Al finalizar el programa se debera visualizar el array con los valores incrementados
	*/

	array := [5]int{4, 2, 5, 6, 7}

	// realizar la funcionalidad
	for i := 0; i < len(array); i++ {
		array[i] += 20
	}

	fmt.Println("Los valores del array son: ", array)
}
