package main

import "fmt"

type MaxHeap struct {
	// left index: 2n+1, right index: 2n+2, parent index: (n-1)/2
	arr []int
}

func (m *MaxHeap) Insert(v int) {
	index := len(m.arr)
	m.arr = append(m.arr, v)
	m.shiftUp(index)
}

func (m *MaxHeap) shiftUp(index int) {
	for m.arr[(index-1)/2] < m.arr[index] {
		m.arr[(index-1)/2], m.arr[index] = m.arr[index], m.arr[(index-1)/2]
		index = (index - 1) / 2
	}
}

func Heapify(arr []int) MaxHeap {
	m := MaxHeap{arr: make([]int, len(arr))}
	for i := 0; i < len(arr); i++ {
		m.arr[i] = arr[i]
	}
	for i := (len(arr) - 1) / 2; i >= 0; i-- {
		m.shiftDown(i)
	}
	return m
}

func (m *MaxHeap) Pop() int {
	v := m.arr[0]
	n := len(m.arr) - 1
	m.arr[0] = m.arr[n]
	index := 0
	m.shiftDown(index)
	return v
}

func (m *MaxHeap) shiftDown(index int) {
	for 2*index+1 <= len(m.arr)-1 {
		var maxIndex int
		if 2*index+2 <= len(m.arr)-1 && m.arr[2*index+1] < m.arr[2*index+2] {
			maxIndex = 2*index + 2
		} else {
			maxIndex = 2*index + 1
		}
		if m.arr[index] < m.arr[maxIndex] {
			m.arr[index], m.arr[maxIndex] = m.arr[maxIndex], m.arr[index]
			index = maxIndex
		} else {
			break
		}
	}
	m.arr = m.arr[:len(m.arr)]
}

func Sort(arr []int, incr bool) []int {
	//m := MaxHeap{}
	//for _, v := range arr {
	//	m.Insert(v)
	//}
	//if incr == true {
	//	for i := len(arr) - 1; i >= 0; i-- {
	//		arr[i] = m.Pop()
	//	}
	//} else {
	//	for i := 0; i < len(arr); i++ {
	//		arr[i] = m.Pop()
	//	}
	//}
	//return arr
	//---------------------------
	//heap := Heapify(arr)
	//if incr == true {
	//	for i := len(arr) - 1; i >= 0; i-- {
	//		arr[i] = heap.Pop()
	//	}
	//} else {
	//	for i := 0; i < len(arr); i++ {
	//		arr[i] = heap.Pop()
	//	}
	//}
	//return arr
	//---------------------------
	n := len(arr)
	for i := (n - 1) / 2; i >= 0; i-- {
		shiftDown(arr, n, i)
	}
	if incr == true {
		for i := n - 1; i >= 0; i-- {
			arr[i], arr[0] = arr[0], arr[i]
			shiftDown(arr, i, 0)
		}
	}

	return arr
}

func shiftDown(arr []int, n, index int) {
	for 2*index+1 <= n-1 {
		var maxIndex int
		if 2*index+2 <= n-1 && arr[2*index+1] < arr[2*index+2] {
			maxIndex = 2*index + 2
		} else {
			maxIndex = 2*index + 1
		}
		if arr[index] < arr[maxIndex] {
			arr[index], arr[maxIndex] = arr[maxIndex], arr[index]
			index = maxIndex
		} else {
			break
		}
	}
	arr = arr[:n]
}

func findKthLargest(arr []int, k int) int {
	n := len(arr)
	for i := (n - 1) / 2; i >= 0; i-- {
		shiftDown(arr, n, i)
	}
	for i := n - 1; i >= n-k+1; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		shiftDown(arr, i, 0)
	}
	return arr[0]
}

func main() {
	//m := MaxHeap{}
	//m.Insert(1)
	//m.Insert(2)
	//m.Insert(3)
	//m.Insert(4)
	//m.Insert(5)
	//fmt.Println(m.arr)
	//fmt.Println(m.Pop())
	//fmt.Println(m.arr)
	//fmt.Println(m.Pop())
	//fmt.Println(m.arr)
	fmt.Println(Sort([]int{3, 5, 2, 6, 1, 4}, true))
	fmt.Println(Sort([]int{3, 5, 2, 6, 1, 4}, false))
	heap := Heapify([]int{3, 5, 2, 6, 1, 4})
	fmt.Println(heap.arr)
	fmt.Println(findKthLargest([]int{3, 5, 2, 6, 1, 4}, 3))
}
