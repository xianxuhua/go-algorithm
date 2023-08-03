package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	n1, n2 := len(a), len(b)
	if n1 > n2 {
		t := ""
		for i := 0; i < n1-n2; i++ {
			t += "0"
		}
		b += t
	} else {
		t := ""
		for i := 0; i < n2-n1; i++ {
			t += "0"
		}
		a += t
	}
	cs := []int{}
	for i := 0; i < len(a); i++ {
		ai, _ := strconv.Atoi(strconv.Itoa(int(a[i])))
		bi, _ := strconv.Atoi(strconv.Itoa(int(b[i])))
		cs = append(cs, ai+bi)
		fmt.Println(ai + bi - '0')
	}

	return ""
}

func main() {
	fmt.Println(addBinary("111", "1"))
}
