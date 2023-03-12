package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	maxA := float64(j-i) * math.Min(float64(height[i]), float64(height[j]))
	for i < j {
		newIA := float64(j-(i+1)) * math.Min(float64(height[i+1]), float64(height[j]))
		newJA := float64((j-1)-i) * math.Min(float64(height[i]), float64(height[j-1]))
		if newIA > newJA {
			i++
		} else {
			j--
		}
		if newIA > maxA {
			maxA = newIA
			i++
		} else if newJA > maxA {
			maxA = newJA
			j--
		}
		fmt.Println(i, j)
	}

	return int(maxA)
}

func main() {
	fmt.Println(maxArea([]int{2, 3, 10, 5, 7, 8, 9}))
}
