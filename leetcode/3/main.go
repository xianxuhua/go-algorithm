package main

import "fmt"

func IsRepeat(s string) bool {
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++
		if m[string(v)] > 1 {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			subStr := s[i : j+1]
			if !IsRepeat(subStr) && len(subStr) > maxLen {
				maxLen = len(subStr)
			}
		}
	}
	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	maxLen := 0
	i, j := 0, -1
	m := make(map[string]int)
	for i < len(s) {
		if j+1 < len(s) && m[string(s[j+1])] < 1 {
			j++
			m[string(s[j])]++
		} else {
			delete(m, string(s[i]))
			i++
		}
		if len(m) > maxLen {
			maxLen = len(m)
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring2("bbbbb"))
}
