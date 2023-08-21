package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	findPermute(nums, 0, []int{}, used, func(ints []int) {
		res = append(res, ints)
	})
	return res
}

func findPermute(nums []int, index int, tmp []int, used []bool, opt func([]int)) {
	if index == len(nums) {
		opt(tmp)
		return
	}

	for i, v := range nums {
		if !used[i] {
			used[i] = true
			findPermute(nums, index+1, append(tmp, v), used, opt)
			used[i] = false
		}
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{5, 4, 6, 2}))
}
