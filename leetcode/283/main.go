package main

import (
	"algorithm/queue"
	"fmt"
)

func moveZeros(arr []int) []int {
	res := make([]int, len(arr))
	q := queue.Queue{}
	for _, v := range arr {
		if v != 0 {
			q.Push(v)
		}
	}
	n := q.Len()
	for i := 0; i < n; i++ {
		res[i] = q.Pop()
	}
	return res
}

func moveZeros2(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[j] = arr[i]
			j++
		}
	}
	for i := j; i < len(arr); i++ {
		arr[i] = 0
	}
	return arr
}

func moveZeros3(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}

	return arr
}

func moveZeros4(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			if i != j {
				arr[j], arr[i] = arr[i], arr[j]
				j++
			} else {
				j++
			}
		}
	}

	return arr
}

func main() {
	//target := []int{1, 3, 12, 0, 0}
	fmt.Println(moveZeros4([]int{0, 1, 0, 3, 12}))
	fmt.Println(moveZeros4([]int{0, 1, 0, 3, 12}))
}
