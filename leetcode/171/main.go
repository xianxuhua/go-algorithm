package main

import "fmt"

func titleToNumber(columnTitle string) int {
	r := 0
	for i, v := range columnTitle {

		r += int(v-64) * (len(columnTitle) - i - 1) * 10
	}
	return r
}

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("Z"))
	fmt.Println(titleToNumber("AA"))
	fmt.Println(titleToNumber("AB"))
}
