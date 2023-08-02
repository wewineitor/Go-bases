package main

import (
	"fmt"

	"github.com/wewineitor/Go-bases/06-funciones/function"
)

func main() {
	res := function.Add(3, 4)
	fmt.Println(res)

	v, err := function.Calc(function.SUM, 3, 6)
	fmt.Println("Value: ", v, " - Error: ", err)

	v, err = function.Calc(6, 3, 6)
	fmt.Println("Value: ", v, " - Error: ", err)

	v = function.MSum(23, 12, 32, 3, 6, 10, 70)
	fmt.Println("multy sum ", v)
}
