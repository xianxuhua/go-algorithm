package main

import "fmt"

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	//defer修改返回值
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}

func main() {
	fmt.Println(b())
}
