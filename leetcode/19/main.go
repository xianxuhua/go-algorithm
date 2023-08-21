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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	pre := &ListNode{}
	dummyHead := pre
	pre.Next = head
	cur = head
	count2 := 0
	for cur != nil {
		count2++
		if count2 == count-n+1 {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}
	return dummyHead.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	p := dummyHead
	q := dummyHead
	for i := 0; i < n+1; i++ {
		q = q.Next
	}
	for q != nil {
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return dummyHead.Next
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
	printLinkedList(removeNthFromEnd2(l, 2))
	printLinkedList(removeNthFromEnd2(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}, 2))
}
