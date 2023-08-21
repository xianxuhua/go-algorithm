package main

import "fmt"

func max3(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else {
		return c
	}
}

func integerBreak(n int) int {
	cache := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		cache[i] = 1
	}
	for i := 2; i <= n; i++ {
		// 尝试对i进行分割， j + (i-j)
		for j := 1; j <= i-1; j++ {
			cache[i] = max3(cache[i-j]*j, (i-j)*j, cache[i])
		}
	}
	return cache[n]
}

func integerBreak2(n int) int {
	if n == 1 {
		return 1
	}

	res := -1
	for i := 1; i < n; i++ {
		res = max3(integerBreak2(n-i)*i, (n-i)*i, res)
	}
	return res
}

func main() {
	fmt.Println(integerBreak(30) == integerBreak2(30))
}
