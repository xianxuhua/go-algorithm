package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
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
		findCombine(candidates, target, i, sum+candidates[i], append(tmp, candidates[i]), opt)
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
