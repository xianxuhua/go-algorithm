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

// [l...r)
func binarySearch2(arr []int, l, r, target int) int {
	if l >= r {
		return -1
	}
	mid := (l + r) / 2
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binarySearch2(arr, l, mid, target)
	} else {
		return binarySearch2(arr, mid+1, r, target)
	}
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
