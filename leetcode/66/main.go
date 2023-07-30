package main

import (
	"fmt"
	"strconv"
	"strings"
)

func plusOne(digits []int) []int {
	stringDigits := []string{}
	for _, v := range digits {
		stringDigits = append(stringDigits, strconv.Itoa(v))
	}

	i, _ := strconv.Atoi(strings.Join(stringDigits, ""))
	fmt.Println(i)
	i++
	a := strconv.Itoa(i)
	res := []int{}
	for _, v := range a {
		atoi, _ := strconv.Atoi(string(v))
		res = append(res, atoi)
	}
	return res
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
}
