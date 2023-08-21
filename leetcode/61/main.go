package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printLinkedList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, "->")
		cur = cur.Next
	}
	fmt.Println()
}

func rotateRight(head *ListNode, k int) *ListNode {
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	p, q := head, head
	for i := 0; i < k%count; i++ {
		q = q.Next
	}
	for q.Next != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = head
	ret := p.Next
	p.Next = nil
	return ret
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	printLinkedList(rotateRight(l, 2))
}
