package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1, m2 := make(map[int]int), make(map[int]int)
	for i := 0; i < len(s); i++ {
		m1[int(s[i])]++
		m2[int(t[i])]++
	}
	for k, _ := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
