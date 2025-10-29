package math

import (
	"errors"
)

func Add(a int, b int) []int {
	k := make([]int, 1, a+b)
	k[0] = a + b
	// k := []int{a + b}
	return k
}

func Division(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("integer divison by Zero")
	}
	return a / b, nil
}

// variable lenght arrays can be passed.
func SquareArray(arr *[]float64) []float64 {
	for a := range *arr {
		(*arr)[a] = (*arr)[a] * (*arr)[a]
	}
	return *arr
}
