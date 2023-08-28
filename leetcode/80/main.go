package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 0
	j := len(nums) - 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
			if count >= 2 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		} else {
			j--
			count = 0
		}
	}
	fmt.Println(nums, j, nums[j:])
	return len(nums) - j
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
	// 1, 1, 1, 2, 2, 2, 3
	// j i-1 i
}
