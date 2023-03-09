package main

import "fmt"

func sortColors(nums []int) {
	i, j, k := 0, 0, 0
	for m := 0; m < len(nums); m++ {
		if nums[m] == 0 {
			i++
		} else if nums[m] == 1 {
			j++
		} else {
			k++
		}
	}

	for m := 0; m < i; m++ {
		nums[m] = 0
	}
	for m := i; m < i+j; m++ {
		nums[m] = 1
	}
	for m := i + j; m < i+j+k; m++ {
		nums[m] = 2
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
