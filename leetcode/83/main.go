package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for cur != nil {
		cur = cur.Next
	}
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	fmt.Println(deleteDuplicates(l))
}
