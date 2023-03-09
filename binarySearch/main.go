package main

import (
	"fmt"
	"time"
)

// ****
func BinarySearch(arr []int, v int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] > v {
			r = mid - 1
		} else if arr[mid] < v {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{}
	n := 10000000
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}

	startTime := time.Now()
	for i := 0; i < n; i++ {
		r := BinarySearch(arr, i)
		if i != r {
			panic("")
		}
	}

	fmt.Println(time.Since(startTime))

}
