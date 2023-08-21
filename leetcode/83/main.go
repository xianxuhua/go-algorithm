package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var pre = &ListNode{
		Val:  -1,
		Next: head,
	}
	cur := pre.Next
	newHead := pre.Next

	for cur != nil {
		if pre.Val == cur.Val {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}
	return newHead
}

func printLinkedList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, "->")
		cur = cur.Next
	}
	fmt.Println()
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	printLinkedList(deleteDuplicates(l))
}
