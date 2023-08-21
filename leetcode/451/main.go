package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	m := make(map[int]int)
	for _, v := range s {
		m[int(v)]++
	}
	ks := []int{}
	for k := range m {
		ks = append(ks, k)
	}
	sort.SliceStable(ks, func(i, j int) bool {
		return m[ks[i]] < m[ks[j]]
	})
	res := ""
	for i := len(ks) - 1; i >= 0; i-- {
		for j := 0; j < m[ks[i]]; j++ {
			res += fmt.Sprintf("%c", ks[i])
		}
	}

	return res
}

func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
}
