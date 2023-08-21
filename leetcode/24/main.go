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

func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	newDummyHead := pre
	// 先确保pre.next != nil
	for pre.Next != nil && pre.Next.Next != nil {
		node1 := pre.Next
		node2 := node1.Next
		next := node2.Next

		node2.Next = node1
		node1.Next = next
		pre.Next = node2

		pre = node1

	}
	return newDummyHead.Next
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
				},
			},
		},
	}
	printLinkedList(swapPairs(l))
}
