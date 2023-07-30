package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for a != nil {
		for b != nil {
			if a == b {
				return a
			}
			b = b.Next
		}
		b = headB
		a = a.Next
	}
	return nil
}

func main() {
	l1Last := &ListNode{
		Val: 1,
	}
	l1 := &ListNode{
		Val:  4,
		Next: l1Last,
	}
	l2Last := &ListNode{
		Val: 1,
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  6,
			Next: l2Last,
		},
	}
	com := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	l1Last.Next = com
	l2Last.Next = com
	fmt.Println(getIntersectionNode(l1, l2).Val)
}
