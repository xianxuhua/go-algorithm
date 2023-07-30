package main

import (
	"fmt"
	"math"
	"strconv"
)

func isHappy(n int) bool {
	l := len(strconv.Itoa(n))
	for {
		r := 0
		for i := 0; i < l; i++ {
			v := n % 10
			n = n / 10
			r += int(math.Pow(float64(v), 2))
		}
		n = r
		l = len(strconv.Itoa(n))
		if n == 1 {
			return true
		}
	}
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}
