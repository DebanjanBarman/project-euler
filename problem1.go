package main

import (
	"fmt"
)

// GetSumOfTwoMultiPles returns the sum of multiples of a, b with a range r where a, b, r > 0. Example,
// If we list all the natural numbers below 10 that are multiples of 3 or 5
//  we get 3, 5, 6 and 9. The sum of these multiples is 23.
func GetSumOfTwoMultiPles(a, b, limit int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		if i%a == 0 {
			sum += i
		} else if i%b == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(GetSumOfTwoMultiPles(3, 5, 1000))
}
