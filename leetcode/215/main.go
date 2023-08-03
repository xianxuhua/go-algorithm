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

func shiftDown(arr []int, n, index int) {
	for 2*index+1 <= n-1 {
		var maxIndex int
		if 2*index+2 <= n-1 && arr[2*index+1] < arr[2*index+2] {
			maxIndex = 2*index + 2
		} else {
			maxIndex = 2*index + 1
		}
		if arr[index] < arr[maxIndex] {
			arr[index], arr[maxIndex] = arr[maxIndex], arr[index]
			index = maxIndex
		} else {
			break
		}
	}
	arr = arr[:n]
}

func findKthLargest2(arr []int, k int) int {
	n := len(arr)
	for i := (n - 1) / 2; i >= 0; i-- {
		shiftDown(arr, n, i)
	}
	for i := n - 1; i >= n-k+1; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		shiftDown(arr, i, 0)
	}
	return arr[0]
}

func main() {
	fmt.Println(findKthLargest2([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
