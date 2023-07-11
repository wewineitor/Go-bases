package main

import (
	"fmt"
)

func main() {
	var intVar int = 12
	fmt.Println(intVar)

	//Las variables de tipo uint solo pueden ser positivas
	var uintVar uint = 10
	fmt.Println(uintVar)

	var stringVar string = "string"
	fmt.Println(stringVar)

	var boolVar bool = true
	fmt.Println(boolVar)

	//Al darle valor con := se instancia la variable y le asigna un tipo de dato dependiendo del valor
	instancedVariable := "asd"
	fmt.Println(instancedVariable)

	const constVariable int = 20
	fmt.Println(constVariable)

	var A byte = 'A'
	var R byte = 82
	arr := []byte{69, 68, 87, 73, 78}
	fmt.Println(A)
	fmt.Println(R)
	fmt.Println(arr)
	fmt.Println(string(A))
	fmt.Println(string(R))
	fmt.Println(string(R))
	fmt.Println(string(arr))
}
