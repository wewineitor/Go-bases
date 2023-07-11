package main

import (
	"fmt"
)

func main() {
	var vectorInt [3]int
	vectorString := [3]string{"perro", "oso", "gato"}

	vectorInt[0] = 10
	vectorInt[1] = 20
	vectorInt[2] = 30

	fmt.Println(vectorInt)
	fmt.Println(vectorString)

	var slice1 []int
	slice1 = append(slice1, 10, 20, 30, 40)
	fmt.Println(slice1)
	fmt.Println(slice1[0:2])
	fmt.Println(len(slice1))

	sliceInstanced := make([]int, 3)
	fmt.Println(sliceInstanced)
}
