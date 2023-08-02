package function

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Add(x int, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Println(value)
	}
}

// Funcion que retorna 2 valores
func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("y must'n be zero")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	}
	return 0, errors.New("has been an error")
}

func MSum(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}

	return sum
}
