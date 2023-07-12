package main

import "fmt"

// 编译器优先进行逃逸分析，逃逸的才分配到堆上
// i和map1是分配在堆上的，而j和map2是在栈上的
func test() *int {
	var i *int = new(int)
	var j *int = new(int)
	println(*i, *j)
	return i
}

func test2() map[int]int {
	map1 := make(map[int]int, 1)
	map2 := make(map[int]int, 1)
	map1[1] = 1
	map2[1] = 1
	return map1
}

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
