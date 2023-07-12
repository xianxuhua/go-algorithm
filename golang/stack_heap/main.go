package main

import "fmt"

func Add() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}

func main() {
	fn := Add()
	fmt.Println(fn()) // 1
	fmt.Println(fn()) // 2
	//	Add() 返回的是一个闭包，并且该闭包访问了外部变量num，那么num将会被分配到堆上
}
