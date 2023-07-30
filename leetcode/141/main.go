package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	m := make(map[string]int)
	cur := head
	for cur != nil {
		q := fmt.Sprintf("%q", cur)
		if _, ok := m[q]; !ok {
			m[q]++
			cur = cur.Next
		} else {
			return true
		}
	}
	return false
}

func main() {
	lastNode := &ListNode{
		Val: 2,
	}
	//circleNode := &ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val:  0,
	//		Next: lastNode,
	//	},
	//}
	l := &ListNode{
		Val:  3,
		Next: lastNode,
	}
	lastNode.Next = l

	fmt.Println(hasCycle(l))
}
