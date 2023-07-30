package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
