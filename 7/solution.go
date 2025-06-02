package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var isNegative bool = x < 0
	if isNegative {
		x *= -1
	}

	var reversed int
	for x > 0 {
		reversed = reversed*10 + (x % 10)
		x = x / 10
	}

	if reversed >= math.MaxInt32 {
		return 0
	}

	if isNegative {
		reversed *= -1
	}

	return reversed
}

func main() {
	fmt.Println(reverse(1534236469))
}
