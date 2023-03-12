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
	//容量小于 1024时，使用2倍扩容，大于1024时，使用1.76倍扩容
	for i := 0; i < 1000; i++ {
		a2 = append(a2, i)
		fmt.Println(cap(a2), len(a2))
	}
}
