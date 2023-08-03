package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i]+nums[j] == target {
	//			return []int{i, j}
	//		}
	//	}
	//}
	//return []int{0, 0}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; !ok {
			m[nums[i]] = i
		} else {
			return []int{v, i}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{1, 4, 1, 1, 2}, 6))
}
