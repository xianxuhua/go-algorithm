package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	First  int
	Second int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	pairs := []Pair{}
	for k, v := range m {
		pairs = append(pairs, Pair{First: k, Second: v})
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].Second > pairs[j].Second
	})
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, pairs[i].First)
	}
	return res
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 3))
}
