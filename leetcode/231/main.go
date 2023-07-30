package main

import (
	"fmt"
	"math"
)

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	log := math.Log2(float64(n))
	return log == int(log)
}

func main() {
	fmt.Println(isPowerOfTwo(2))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(3))
}
