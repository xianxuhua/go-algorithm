package main

import "fmt"

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
