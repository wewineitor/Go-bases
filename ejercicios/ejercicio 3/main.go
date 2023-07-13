package main

import (
	"fmt"
)

func main() {
	codigos := map[int]string{
		10: "notebook",
		15: "tv",
		21: "heladera",
		27: "monitor",
		35: "camara",
	}
	var codigo int = 10
	var array []string

	for codigo != 0 {
		fmt.Print(":")
		fmt.Scanln(&codigo)

		if codigos[codigo] != "" {
			array = append(array, codigos[codigo])
		} else {
			if codigo != 0 {
				array = append(array, "No encontrado")
			}
		}
	}
	fmt.Println(array)
}
