package main

import (
	"fmt"
)

func main() {
	selectionSort([]int{
		8, 6, 2, 3, 1, 5, 7, 4,
	})
	insertionSort([]int{
		8, 6, 2, 3, 1, 5, 7, 4,
	})
	mergeSort([]int{
		8, 6, 2, 3, 1, 5, 7, 4,
	})
	quickSort([]int{
		8, 6, 2, 3, 1, 5, 7, 4,
	})
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	fmt.Println(arr)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		// j代表当前的要交换的范围
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func mergeSort(arr []int) {
	res := _mergeSort(arr)
	fmt.Println(res)
}

func _mergeSort(arr []int) []int {
	midIndex := len(arr) / 2
	if midIndex <= 0 {
		return arr
	}
	return merge(_mergeSort(arr[:midIndex]), _mergeSort(arr[midIndex:]))
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

func quickSort(arr []int) {

}
