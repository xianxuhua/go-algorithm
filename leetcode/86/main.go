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

func partition(head *ListNode, x int) *ListNode {
	var newHeadSmall = &ListNode{}
	var newHeadBig = &ListNode{}
	cur := head
	newSmallCur := newHeadSmall
	newBigCur := newHeadBig
	for cur != nil {
		if cur.Val < x {
			newSmallCur.Next = &ListNode{Val: cur.Val}
			newSmallCur = newSmallCur.Next
		} else {
			newBigCur.Next = &ListNode{Val: cur.Val}
			newBigCur = newBigCur.Next
		}
		cur = cur.Next
	}
	newSmallCur.Next = newHeadBig.Next
	return newHeadSmall.Next
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}
	printLinkedList(partition(l, 3))
}
