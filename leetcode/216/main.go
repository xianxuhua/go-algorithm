package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	find(k, n, 1, 0, []int{}, func(ints []int) {
		res = append(res, ints)
	})
	return res
}

func find(k, n, start, sum int, temp []int, opt func([]int)) {
	if sum > n {
		return
	}

	if sum == n && len(temp) == k {
		newTemp := make([]int, len(temp))
		copy(newTemp, temp)
		opt(newTemp)
		return
	}

	for i := start; i <= 9; i++ {
		find(k, n, i+1, sum+i, append(temp, i), opt)
	}
}

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
	fmt.Println(combinationSum3(9, 45))
}
