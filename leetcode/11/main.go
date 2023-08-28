package main

import (
	"fmt"
)

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for i < j {
		area := (j - i) * min(height[i], height[j])
		res = max(area, res)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return res
}

func main() {
	fmt.Println(maxArea([]int{2, 3, 10, 5, 7, 8, 9}))
}
