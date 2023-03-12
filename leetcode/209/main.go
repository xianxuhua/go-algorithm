package main

import (
	"fmt"
)

func sum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

func minSubArrayLen(target int, nums []int) int {
	minLen := len(nums)
	find := false
	for i := 0; i < len(nums); i++ {
		s := 0
		for j := i; j < len(nums); j++ {
			//s := sum(nums[i : j+1])
			s += nums[j]
			if s >= target && j-i+1 <= minLen {
				minLen = j - i + 1
				find = true
			}
		}
	}
	if !find {
		return 0
	}
	return minLen
}

// 难
func minSubArrayLen2(target int, nums []int) int {
	i, j := 0, -1
	s := 0
	find := false
	minLen := len(nums)
	for i < len(nums) {
		if s < target && j+1 < len(nums) {
			j++
			s += nums[j]
		} else {
			s -= nums[i]
			i++
		}
		//if s >= target {
		//
		//}
		if s >= target && j-i+1 <= minLen {
			minLen = j - i + 1
			find = true
		}

	}
	if !find {
		return 0
	}
	return minLen
}

func main() {
	fmt.Println(minSubArrayLen2(6, []int{10, 2, 3}))
}
