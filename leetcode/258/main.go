package main

import (
	"fmt"
	"strconv"
)

func addDigits(num int) int {
	n := len(strconv.Itoa(num))
	for n != 1 {
		t := 0
		for i := 0; i < n; i++ {
			t += num % 10
			num = num / 10
		}
		n = len(strconv.Itoa(t))
		num = t
	}

	return num
}

func main() {
	fmt.Println(addDigits(38))
	fmt.Println(addDigits(0))
	fmt.Println(addDigits(99))
}
