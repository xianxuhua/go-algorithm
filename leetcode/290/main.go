package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	m := make(map[int]string)
	split := strings.Split(s, " ")
	for i := 0; i < len(pattern); i++ {
		v, ok := m[int(pattern[i])]
		fmt.Println(ok, pattern[i], v, split[i])
		if !ok && v != split[i] {
			m[int(pattern[i])] = split[i]
			continue
		} else {
			if v != split[i] {
				return false
			}

		}
	}

	return true
}

func main() {
	//fmt.Println(wordPattern("abba", "dog cat cat dog"))
	//fmt.Println(wordPattern("abba", "dog cat cat fish"))
	//fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}
