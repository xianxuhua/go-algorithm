package main

import "fmt"

func change(arr []int) {
	arr[0] = 1
}
func main() {
	arr := []int{2}
	change(arr)
	fmt.Println(arr)
}
