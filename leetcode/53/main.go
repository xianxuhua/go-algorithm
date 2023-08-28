package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	res := math.MinInt
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			res = max(res, sum)
		}
	}
	return res
}

func maxSubArray2(nums []int) int {
	cache := make([]int, len(nums))
	cache[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		cache[i] = max(cache[i-1]+nums[i], nums[i])
	}
	res := math.MinInt
	for i := 0; i < len(cache); i++ {
		res = max(res, cache[i])
	}
	return res
}

func main() {
	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray2([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray2([]int{-2, -1}))
	fmt.Println(maxSubArray2([]int{-1}))
}
