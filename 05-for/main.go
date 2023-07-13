package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum++
	}

	fmt.Println(sum)

	sum = 1

	for sum < 100 {
		sum++
	}

	fmt.Println(sum)

	sum = 0
	for {
		if sum > 1000 {
			break
		}
		sum++
	}
	fmt.Println(sum)

	arr := []int{1, 2, 3, 4, 5}

	for i := range arr {
		fmt.Println(arr[i])
	}

	map1 := map[string]float64{
		"A": 12.3,
		"B": 32.45,
	}

	for key, value := range map1 {
		fmt.Println("Key:", key, "Value:", value)
	}
}
