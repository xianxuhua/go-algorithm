package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	m1 := make(map[uint8]uint8)
	m2 := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		if _, ok := m1[s[i]]; !ok {
			m1[s[i]] = t[i]
		}
		if _, ok := m2[t[i]]; !ok {
			m2[t[i]] = s[i]
		}
		if t[i] != m1[s[i]] {
			return false
		}
		if s[i] != m2[t[i]] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("a", "a"))
	fmt.Println(isIsomorphic("ab", "ac"))
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
}
