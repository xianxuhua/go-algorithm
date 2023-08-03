package main

import "fmt"

func exists(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return true
		}
	}

	return false
}

func removeDuplicates(nums []int) int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		find := exists(res, nums[i])
		if !find {
			res = append(res, nums[i])
		}
	}

	copy(nums, res)
	return len(nums[:len(res)])
}

func removeDuplicates1(nums []int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[i], nums[k] = nums[k], nums[i]
		}
	}
	return k + 1
}

func main() {
	fmt.Println(removeDuplicates1([]int{1, 1, 2}) == 2)
}
