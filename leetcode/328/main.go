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

func oddEvenList(head *ListNode) *ListNode {
	newHeadOdd := &ListNode{}
	newHeadEven := &ListNode{}
	newCurOdd := newHeadOdd
	newCurEven := newHeadEven
	cur := head
	index := 0
	for cur != nil {
		index++
		if index%2 == 1 {
			newCurOdd.Next = &ListNode{Val: cur.Val}
			newCurOdd = newCurOdd.Next
		} else {
			newCurEven.Next = &ListNode{Val: cur.Val}
			newCurEven = newCurEven.Next
		}
		cur = cur.Next
	}
	newCurOdd.Next = newHeadEven.Next
	return newHeadOdd.Next
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
	printLinkedList(oddEvenList(l))
}
