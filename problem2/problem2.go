package main

import (
	"fmt"
)

// SumOfEvenFibonacciValues returns the sum of even fibonacci values, in a range r where r > 0. Example,
// while range is 40 terms will be: 1, 2, 3, 5, 8, 13, 21, 34
// and the sum of even values will be 2+8+34 = 44
func SumOfEvenFibonacciValues(limit int) int {
	sum := 0
	i := 1
	j := 1

	for {
		c := i + j
		i = j
		j = c

		if c > limit {
			break
		}

		if c%2 == 0 {
			sum += c
		}

	}

	return sum
}

func main() {
	fmt.Println(SumOfEvenFibonacciValues(4000000))
}
