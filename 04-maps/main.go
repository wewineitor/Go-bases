package main

import (
	"fmt"
)

func main() {
	map1 := make(map[int]string)

	map1[10] = "A"
	map1[20] = "B"
	map1[30] = "C"
	fmt.Println(map1)
	fmt.Println(map1[20])

	delete(map1, 20)
	fmt.Println(map1)

	map1[10] = "A2"
	fmt.Println(map1)

	map2 := map[int]string{
		1: "A",
		2: "B",
	}

	fmt.Println(map2)
}
