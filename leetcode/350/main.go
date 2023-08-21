package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]]++
	}
	res := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		v1, ok1 := m1[nums2[i]]
		v2, _ := m2[nums2[i]]
		if ok1 && (v1 == v2 || v1 < v2 && res[nums2[i]] < v1 || v2 < v1 && res[nums2[i]] < v2) {
			res[nums2[i]]++
		}
	}
	arr := []int{}
	for k, v := range res {
		for i := 0; i < v; i++ {
			arr = append(arr, k)
		}
	}
	return arr
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5, 9, 9}, []int{9, 4, 9, 8, 4}))
}
