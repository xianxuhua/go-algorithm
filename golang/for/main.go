package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := make([]*int, len(arr1))

	for i, v := range arr1 {
		// for 循环中使用循环变量的地址，导致前面的v被覆盖
		arr2[i] = &v
	}

	for _, v := range arr2 {
		fmt.Println(*v)
	}
}
