package queue

import "fmt"

type Queue struct {
	arr []int
}

func (q *Queue) Push(e int) {
	q.arr = append(q.arr, e)
}
func (q *Queue) Pop() int {
	v := q.arr[0]
	q.arr = q.arr[1:]
	return v
}
func (q *Queue) Len() int {
	return len(q.arr)
}

func main() {
	q := Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
