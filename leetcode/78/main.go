package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{}
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
		find(nums, i+1, append(tmp, nums[i]), opt)
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3, 4}))
}
