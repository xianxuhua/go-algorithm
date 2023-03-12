package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		t := target - numbers[i]
		l, r := i+1, len(numbers)-1
		for l <= r {
			mid := (l + r) / 2
			if numbers[mid] > t {
				r = mid - 1
			} else if numbers[mid] < t {
				l = mid + 1
			} else {
				return []int{i, mid}
			}
		}
	}
	return []int{}
}

func twoSum2(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i <= j {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			return []int{i, j}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum2([]int{2, 7, 11, 15}, 9))
}
