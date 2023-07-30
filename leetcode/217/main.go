package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for _, v := range m {
		if v >= 2 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
