package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	newArr := make([]int, len(nums1))
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			newArr[k] = nums1[i]
			i++
			k++
		} else if nums1[i] > nums2[j] {
			newArr[k] = nums2[j]
			j++
			k++
		}
	}
	for i < m {
		newArr[k] = nums1[i]
		i++
		k++
	}
	for j < n {
		newArr[k] = nums2[j]
		j++
		k++
	}
	copy(nums1, newArr)
	fmt.Println(nums1)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{1}, 1, []int{}, 0)
	merge([]int{0}, 0, []int{1}, 1)
}
