package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	m1 := make(map[uint8]string)
	m2 := make(map[string]uint8)
	split := strings.Split(s, " ")
	if len(pattern) != len(split) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if _, ok := m1[pattern[i]]; !ok {
			m1[pattern[i]] = split[i]
		}
		if _, ok := m2[split[i]]; !ok {
			m2[split[i]] = pattern[i]
		}
		if split[i] != m1[pattern[i]] {
			return false
		}
		if pattern[i] != m2[split[i]] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}
