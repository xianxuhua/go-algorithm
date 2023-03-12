package main

import (
	"fmt"
)

func isVowel(a int32) bool {
	if a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' || a == 'A' || a == 'E' || a == 'I' || a == 'O' || a == 'U' {
		return true
	}
	return false
}

func reverseVowels(s string) string {
	i, j := 0, len(s)-1
	t := []int32{}
	for _, v := range s {
		t = append(t, v)
	}
	for i < j {
		if isVowel(t[i]) && isVowel(t[j]) {
			t[i], t[j] = t[j], t[i]
			i++
			j--
		} else {
			for !isVowel(t[i]) && i < j {
				i++
			}
			for !isVowel(t[j]) && i < j {
				j--
			}
		}

	}
	res := ""
	for _, v := range t {
		res += fmt.Sprintf("%c", v)
	}
	return res
}

func main() {
	fmt.Println(reverseVowels(".,"))
}
