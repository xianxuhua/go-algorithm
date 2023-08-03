package main

import (
	"fmt"
)

func index(arr []int, v int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] > v {
			r = mid - 1
		} else if arr[mid] < v {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func intersection(nums1 []int, nums2 []int) []int {
	// find min in max arr index
	indexes := []int{}
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		// find nums2 in nums1 index
		for i := 0; i < n2; i++ {
			i2 := index(nums1, nums2[i])
			if i2 != -1 {
				indexes = append(indexes, i2)
			}
		}
	} else {
		for i := 0; i < n1; i++ {
			i2 := index(nums2, nums1[i])
			if i2 != -1 {
				indexes = append(indexes, i2)
			}
		}
	}
	fmt.Println(indexes)
	return []int{}
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}
