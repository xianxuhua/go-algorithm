package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	if n == 1 {
		return 1
	}

	res := 0
	for i := 1; i < n; i++ {
		if math.Sqrt(float64(i)) {
			res = i + numSquares(n-i)
		}
	}
}

func main() {
	fmt.Println(numSquares(12))
}
