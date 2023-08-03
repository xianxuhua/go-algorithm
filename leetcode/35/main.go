package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= target && target <= nums[i+1] {
			return i + 1
		}
	}
	return len(nums)
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}
