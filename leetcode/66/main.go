package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	digits[len(digits)-1] += 1
	if digits[len(digits)-1] < 10 {
		return digits
	}
	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] >= 10 {
			digits[i] = 0
			digits[i-1] += 1
		} else {
			break
		}
	}
	if digits[0] >= 10 {
		res := []int{1}
		for i := 0; i < len(digits); i++ {
			res = append(res, 0)
		}
		return res
	} else {
		return digits
	}
}

func main() {
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9}))
	fmt.Println(plusOne([]int{1, 9}))
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
}
