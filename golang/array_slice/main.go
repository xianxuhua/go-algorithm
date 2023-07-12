package main

import "fmt"

func f1(arr [10]int) {
	arr[0] = 1
}

func f2(arr []int) {
	arr[0] = 1
}

func main() {
	//数组是值类型，切片是引用类型
	a1 := [10]int{}
	a2 := make([]int, 10)
	f1(a1)
	f2(a2)
	fmt.Println(a1, a2)
	//	切片一旦扩容，指向一个新的底层数组
	//容量小于 1024时，使用2倍扩容，大于1024时，使用1.43倍扩容
	for i := 0; i < 1200; i++ {
		a2 = append(a2, i)
		fmt.Println(cap(a2), len(a2))
	}

	a3 := make([]int, 2, 5)
	fmt.Println(a3)
}
