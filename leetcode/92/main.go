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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil || left == right {
		return head
	}

	var pre *ListNode = nil
	cur, next := head, head.Next

	saveHead, saveLast := head, head
	for i := 0; i < left-2; i++ {
		saveHead = saveHead.Next
		cur = cur.Next
	}
	for i := 0; i < right; i++ {
		saveLast = saveLast.Next
	}
	cur = cur.Next
	saveHead.Next = nil

	for cur != saveLast {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	preCur := pre
	for preCur.Next != nil {
		preCur = preCur.Next
	}

	head.Next = pre
	preCur.Next = saveLast

	return head
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
	printLinkedList(reverseBetween(l, 2, 4))
	printLinkedList(reverseBetween(&ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}, 1, 1))
}
