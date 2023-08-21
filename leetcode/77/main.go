package main

import "fmt"

func combine(n int, k int) [][]int {
	var res [][]int
	findCombine(n, k, 1, []int{}, func(ints []int) {
		res = append(res, ints)
	})
	return res
}

func findCombine(n, k, start int, temp []int, opt func([]int)) {
	if len(temp) == k {
		newTmp := make([]int, k)
		copy(newTmp, temp)
		opt(newTmp)
		return
	}

	for i := start; i <= n-(k-len(temp))+1; i++ {
		findCombine(n, k, i+1, append(temp, i), opt)
	}
}

func main() {
	fmt.Println(combine(5, 4))
}
