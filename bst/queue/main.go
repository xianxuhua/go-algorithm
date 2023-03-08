package queue

import "algorithm/bst/tree"

type Queue struct {
	arr []*tree.Node
}

func (q *Queue) Push(e *tree.Node) {
	q.arr = append(q.arr, e)
}
func (q *Queue) Pop() *tree.Node {
	v := q.arr[0]
	q.arr = q.arr[1:]
	return v
}

func (q *Queue) Len() int {
	return len(q.arr)
}
