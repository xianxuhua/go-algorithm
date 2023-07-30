package main

import "fmt"

func isUgly(n int) bool {
	if n == 1 {
		return true
	}

	fmt.Println(n % 5)
	return false
}

func main() {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(14))
}
