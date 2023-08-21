package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.SliceStable(g, func(i, j int) bool {
		return g[i] > g[j]
	})
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	gi, si := 0, 0
	res := 0
	for gi < len(g) && si < len(s) {
		if s[si] >= g[gi] {
			res++
			gi++
			si++
		} else {
			gi++
		}
	}
	return res
}

func main() {
	//fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	//fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	fmt.Println(findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8}))
}
