package main

import "fmt"

func summaryRanges(nums []int) []string {
	res := []string{}
	start, end := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			end = nums[i]
		} else {
			fmt.Println(start, end)
			if start == end {
				res = append(res, fmt.Sprintf("%v", start))
			} else {
				res = append(res, fmt.Sprintf("%v->%v", start, end))
			}
			start = nums[i]
		}
	}
	return res
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
