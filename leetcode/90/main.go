package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	find(nums, 0, []int{}, func(ints []int) {
		res = append(res, ints)
	})
	return res
}

func find(nums []int, start int, tmp []int, opt func([]int)) {
	newTmp := make([]int, len(tmp))
	copy(newTmp, tmp)
	opt(newTmp)

	for i := start; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && i > start {
			continue
		}
		find(nums, i+1, append(tmp, nums[i]), opt)
	}
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
