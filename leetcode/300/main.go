package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	cache := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		cache[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				cache[i] = max(cache[i], 1+cache[j])
			}
		}
	}

	res := 1
	for i := 0; i < len(cache); i++ {
		res = max(res, cache[i])
	}
	return res
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
