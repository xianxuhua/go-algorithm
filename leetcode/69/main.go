package main

import (
	"fmt"
)

func mySqrt(x int) int {
	//res := 0
	//for i := 1; i <= x; i++ {
	//	if i*i > x {
	//		break
	//	}
	//	res++
	//}
	//return res
	l, r := 0, x
	res := 0
	for l < r {
		mid := (l + r) / 2

		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid + 1
			res = mid
		}
	}
	return res
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}
