package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	sqrt := math.Sqrt(float64(x))
	return int(sqrt)
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}
