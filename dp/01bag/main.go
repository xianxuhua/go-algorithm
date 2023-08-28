package main

import "fmt"

// 用[0...i]填充容量为c的背包
func bestValue(w, v []int, i, c int) int {
	if i < 0 {
		return 0
	}

	res := bestValue(w, v, i-1, c)
	if c >= w[i] {
		res = max(res, v[i]+bestValue(w, v, i-1, c-w[i]))
	}
	return res
}

func bag01(w, v []int, c int) int {
	n := len(w)
	return bestValue(w, v, n-1, c)
}

func main() {
	fmt.Println(bag01([]int{1, 2, 3}, []int{1, 2, 3}, 4))
}
