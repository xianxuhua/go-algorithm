package main

import "fmt"

func combinationSum2(candidates []int, target int) [][]int {
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

}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
