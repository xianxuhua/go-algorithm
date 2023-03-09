package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	return _findKthLargest(nums, 0, len(nums)-1, len(nums)-k)
}

func _findKthLargest(nums []int, l, r, k int) int {
	p := partition(nums, l, r)
	if p == k {
		return nums[p]
	} else if p <= k {
		return _findKthLargest(nums, p+1, r, k)
	} else {
		return _findKthLargest(nums, 0, p-1, k)
	}
}

func partition(nums []int, l, r int) int {
	randI := rand.Intn(r-l+1) + l
	nums[randI], nums[l] = nums[l], nums[randI]
	j := l
	v := nums[j]
	for i := l + 1; i <= r; i++ {
		if nums[i] <= v {
			nums[i], nums[j+1] = nums[j+1], nums[i]
			j++
		}
	}

	nums[l], nums[j] = nums[j], nums[l]
	return j
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
