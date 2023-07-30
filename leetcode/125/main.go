package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	list := []int{}
	for _, v := range s {
		if v >= 65 && v <= 90 || v >= 97 && v <= 122 || v >= 48 && v <= 57 {
			list = append(list, int(v))
		}
	}
	s2 := ""
	for _, v := range list {
		s2 += fmt.Sprintf("%c", v)
	}
	s2 = strings.ToLower(s2)

	for i := 0; i < len(s2)/2; i++ {
		if s2[i] != s2[len(s2)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("0P"))
}
