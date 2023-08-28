package main

import "fmt"

func threeSum(nums []int) [][]int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if _, ok := m[0-(nums[i]+nums[j])]; ok {
				res = append(res, []int{nums[i], nums[j], 0 - (nums[i] + nums[j])})
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
