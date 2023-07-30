package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) add(val int) {
	cur := l
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}
}

func printList(l *ListNode) {
	cur := l
	for cur != nil {
		fmt.Print(cur.Val, "->")
		cur = cur.Next
	}
	fmt.Println()
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	pre := head
	cur := head.Next

	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}
	if head.Val == val {
		return head.Next
	}
	return head
}

func main() {
	l := &ListNode{Val: 1}
	l.add(1)
	l.add(1)
	l.add(2)
	l.add(2)
	l.add(3)
	printList(l)
	l = removeElements(l, 3)
	printList(l)
}
