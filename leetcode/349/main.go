package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	// set
	s := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		s[nums1[i]]++
	}
	res := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		if _, ok := s[nums2[i]]; ok {
			res[nums2[i]]++
		}
	}
	arr := []int{}
	for k := range res {
		arr = append(arr, k)
	}
	return arr
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}
