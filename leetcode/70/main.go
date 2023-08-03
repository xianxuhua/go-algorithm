package main

import "fmt"

func climbStairs(n int) int {
	//m := make(map[int]int)
	//if n == 1 {
	//	return 1
	//}
	//if n == 2 {
	//	return 2
	//}
	//if _, ok := m[n]; ok {
	//	return m[n]
	//} else {
	//	r := climbStairs(n-1) + climbStairs(n-2)
	//	m[n] = r
	//	return r
	//}

	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {

	fmt.Println(climbStairs(45))
}
