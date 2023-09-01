package heap

type MinHeap struct {
	// left index: 2n+1, right index: 2n+2, parent index: (n-1)/2
	arr []int
}

func (m *MinHeap) Insert(v int) {
	index := len(m.arr)
	m.arr = append(m.arr, v)
	m.shiftUp(index)
}

func (m *MinHeap) shiftUp(index int) {
	for m.arr[(index-1)/2] < m.arr[index] {
		m.arr[(index-1)/2], m.arr[index] = m.arr[index], m.arr[(index-1)/2]
		index = (index - 1) / 2
	}
}

func Heapify(arr []int) MinHeap {
	m := MinHeap{arr: make([]int, len(arr))}
	for i := 0; i < len(arr); i++ {
		m.arr[i] = arr[i]
	}
	for i := (len(arr) - 1) / 2; i >= 0; i-- {
		m.shiftDown(i)
	}
	return m
}

func (m *MinHeap) Pop() int {
	v := m.arr[0]
	n := len(m.arr) - 1
	m.arr[0] = m.arr[n]
	index := 0
	m.shiftDown(index)
	return v
}

func (m *MinHeap) shiftDown(index int) {
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
