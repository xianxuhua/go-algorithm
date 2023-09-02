package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	findCombine(candidates, target, 0, 0, []int{}, func(ints []int) {
		res = append(res, ints)
	})
	return res
}

func findCombine(candidates []int, target int, start int, sum int, tmp []int, opt func([]int)) {
	if sum > target {
		return
	}

	if sum == target {
		newTmp := make([]int, len(tmp))
		copy(newTmp, tmp)
		opt(newTmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		// 同一深度、排序后相邻元素相等，后面的不再考虑
		if i > 0 && candidates[i] == candidates[i-1] && i > start {
			continue
		}
		findCombine(candidates, target, i+1, sum+candidates[i], append(tmp, candidates[i]), opt)
	}
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
