package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 1000000

func generateRandomArr() [n]int {
	rand.NewSource(1)
	var arr [n]int
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

func generateNearOrderedArr(swapNum int) [n]int {
	rand.NewSource(1)
	var arr [n]int
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	for i := 0; i < swapNum; i++ {
		x, y := rand.Intn(n), rand.Intn(n)
		arr[x], arr[y] = arr[y], arr[x]
	}
	return arr
}

func generateManyRepeatArr() [n]int {
	rand.NewSource(1)
	var arr [n]int
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(10)
	}
	return arr
}

func runSort(opt func([]int) []int, arr [n]int) {
	startTime := time.Now()
	target := make([]int, n)
	copy(target, arr[:])
	fmt.Println(target[:10])
	res := opt(target)
	fmt.Println(res[:10])
	fmt.Println(time.Since(startTime))
	fmt.Println("**************")
}

func main() {
	//arr := generateNearOrderedArr(10)
	arr := generateManyRepeatArr()
	//runSort(selectionSort, arr)
	//runSort(insertionSort, arr)
	runSort(mergeSort, arr)
	runSort(quickSort, arr)
}

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		// j代表当前的要交换的范围
		e := arr[i]
		var j int
		for j = i; j > 0; j-- {
			if arr[j-1] > e {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = e
	}
	return arr
}

func mergeSort(arr []int) []int {
	return _mergeSort(arr)
}

func _mergeSort(arr []int) []int {
	midIndex := len(arr) / 2
	//if midIndex <= 0 {
	//	return arr
	//}
	if midIndex <= 5 {
		return insertionSort(arr)
	}

	return merge(_mergeSort(arr[:midIndex]), _mergeSort(arr[midIndex:]))
	//arr1 := _mergeSort(arr[:midIndex])
	//arr2 := _mergeSort(arr[midIndex:])
	//if arr1[len(arr1)-1] > arr2[0] {
	//	return merge(arr1, arr2)
	//} else {
	//	return append(arr1, arr2...)
	//}
}

func merge(arr1, arr2 []int) []int {
	newArr := make([]int, len(arr1)+len(arr2))

	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			newArr[k] = arr1[i]
			i++
			k++
		} else {
			newArr[k] = arr2[j]
			j++
			k++
		}
	}
	for i < len(arr1) {
		newArr[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		newArr[k] = arr2[j]
		j++
		k++
	}
	return newArr
}

func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, l, r int) []int {
	if l >= r {
		return arr
	}

	p := partition(arr, l, r)
	arr = _quickSort(arr, l, p-1)
	arr = _quickSort(arr, p+1, r)
	return arr
}

// arr[l, p-1] < arr[p] < arr[p+1, r]
func partition(arr []int, l, r int) int {
	randI := rand.Intn(r-l) + l
	arr[randI], arr[l] = arr[l], arr[randI]
	j := l
	v := arr[j]
	for i := l + 1; i <= r; i++ {
		if arr[i] <= v {
			arr[i], arr[j+1] = arr[j+1], arr[i]
			j++
		}
	}

	arr[l], arr[j] = arr[j], arr[l]
	return j
}
