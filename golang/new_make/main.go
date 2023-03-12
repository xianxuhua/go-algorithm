package main

import "fmt"

func main() {
	// 返回地址
	a := new(int)
	b := new(string)
	c := new([10]int)

	// 返回变量
	d := make([]int, 10)
	e := make(map[string]int)
	f := make(chan int)
	fmt.Println(*a, *b, *c, d, e, f)
}
