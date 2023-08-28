package main

import "fmt"

func rob(nums []int) int {
	return rob_(nums, 0)
}

func rob_(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	res := 0
	for i := start; i < len(nums); i++ {
		res = max(res, nums[i]+rob_(nums, i+2))
	}
	return res
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
