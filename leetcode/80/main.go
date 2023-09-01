package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	first, second int
}

func removeDuplicates(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; !ok || ok && v < 2 {
			m[nums[i]]++
		}
	}
	pairs := []Pair{}
	for k, v := range m {
		pairs = append(pairs, Pair{first: k, second: v})
	}
	res := []int{}
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].first < pairs[j].first
	})
	for _, v := range pairs {
		for i := 0; i < v.second; i++ {
			res = append(res, v.first)
		}
	}
	copy(nums, res)
	return len(res)
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
	// 1, 1, 1, 2, 2, 2, 3
	// j i-1 i
}
