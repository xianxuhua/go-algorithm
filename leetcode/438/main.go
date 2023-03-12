package main

import (
	"fmt"
	"os"
)

func isAnagrams(subStr, p string) bool {
	m := make(map[rune]int)
	for _, v := range subStr {
		m[v]++
	}
	for _, v := range p {
		if m[v] >= 1 {
			m[v]--
		}
	}

	for k, _ := range m {
		if m[k] == 0 {
			delete(m, k)
		}
	}
	return len(m) == 0
}

func findAnagrams(s string, p string) []int {

	i, j := 0, len(p)-1

	res := []int{}
	for j < len(s) {
		if isAnagrams(s[i:j+1], p) {
			res = append(res, i)
		}
		i++
		j++
	}
	return res
}

func findAnagrams2(s string, p string) []int {
	m1, m2 := make(map[rune]int), make(map[rune]int)
	i, j := 0, 0
	for _, v := range p {
		m2[v]++
	}

	res := []int{}
	for j < len(s) {
		if m2[rune(s[j])] >= 1 && m1[rune(s[j])] < m2[rune(s[j])] {
			m1[rune(s[j])]++
			j++
		} else if m2[rune(s[j])] == 0 {
			for k := range m1 {
				delete(m1, k)
			}
			j++
			i = j
		} else if m2[rune(s[j])] >= 1 && m1[rune(s[j])] > 0 {
			m1[rune(s[i])]--
			if m1[rune(s[i])] == 0 {
				delete(m1, rune(s[i]))
			}
			i++
		}

		same := true
		for k := range m1 {
			if m1[k] != m2[k] {
				same = false
			}
		}
		if same && len(m1) == len(m2) {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	file, err := os.ReadFile("/Users/xxh/projects/go/algorithm/leetcode/438/file1.txt")
	if err != nil {
		panic(err)
	}
	file2, err := os.ReadFile("/Users/xxh/projects/go/algorithm/leetcode/438/file2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(findAnagrams2(string(file), string(file2)))
}
