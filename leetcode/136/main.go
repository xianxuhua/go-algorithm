package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}
