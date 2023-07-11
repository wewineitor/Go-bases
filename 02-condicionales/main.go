package main

import (
	"fmt"
)

func main() {
	var number int = 18

	if number >= 18 {
		fmt.Println("Mayor")
	} else {
		fmt.Println("Menor")
	}

	switch number {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}

}
