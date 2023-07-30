package main

import "fmt"

//func removeElement(nums []int, val int) int {
//	j := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == val {
//			nums[i], nums[j] = nums[j], nums[i]
//			j++
//		}
//	}
//	res := len(nums) - j
//	for i := 0; i < res; i++ {
//		nums[i] = nums[j]
//		j++
//	}
//
//	return res
//}

func removeElement(nums []int, val int) int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			res = append(res, nums[i])
		}
	}

	copy(nums, res)
	return len(nums[:len(res)])
}

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
