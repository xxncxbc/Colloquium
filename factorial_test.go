package main

import (
	"fmt"
	"testing"
)

func sliceEq(a, b []int64) bool {
	if a == nil && b == nil {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestFactorial(t *testing.T) {

	var tests = []struct {
		n       int64
		want    int64
		wanterr error
	}{
		{0, 1, nil},
		{1, 1, nil},
		{12, 479001600, nil},
		{9, 362880, nil},
		{-1, -1, errNegativeInputFactorial},
		{-199, -1, errNegativeInputFactorial},
	}

	for _, test := range tests {
		name := fmt.Sprintf("case(%d)", test.n)
		t.Run(name, func(t *testing.T) {
			got, goterr := Factorial(test.n)
			if got != test.want || goterr != goterr {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}

func TestArrFactorial(t *testing.T) {
	var tests = []struct {
		n       int64
		want    []int64
		wanterr error
	}{
		{0, []int64{1}, nil},
		{1, []int64{1, 1}, nil},
		{-1, nil, errNegativeInputArrFactorial},
		{6, []int64{1, 1, 2, 6, 24, 120, 720}, nil},
	}

	for _, test := range tests {
		name := fmt.Sprintf("case(%d)", test.n)
		t.Run(name, func(t *testing.T) {
			got, goterr := ArrFactorial(test.n)
			if goterr != test.wanterr || !sliceEq(got, test.want) {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
