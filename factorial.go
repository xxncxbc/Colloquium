package main

import (
	"errors"
	"fmt"
)

var errNegativeInputFactorial = errors.New("number is negative, unable to calculate factorial")

var errNegativeInputArrFactorial = errors.New("cannot create slice, number is negative")

func Factorial(n int64) (int64, error) {
	if n < 0 {
		return -1, errNegativeInputFactorial
	}
	var res int64 = 1
	for ; n > 0; n-- {
		res *= n
	}
	return res, nil
}

func ArrFactorial(n int64) ([]int64, error) {
	if n < 0 {
		return nil, errNegativeInputArrFactorial
	}
	res := make([]int64, n+1)
	var i int64 = 0
	for ; i < n+1; i++ {
		temp, err := Factorial(i)
		if err == nil {
			res[i] = temp
		}
	}
	return res, nil
}

func main() {
	fmt.Println(ArrFactorial(-1))
}
