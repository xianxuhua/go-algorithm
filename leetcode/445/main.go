package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var pre *ListNode = nil
	cur, next := head, head.Next
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func printLinkedList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, "->")
		cur = cur.Next
	}
	fmt.Println()
}

func Len(l *ListNode) int {
	res := 0
	cur := l
	for cur != nil {
		res++
		cur = cur.Next
	}
	return res
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
	n := min(Len(l1), Len(l2))
	dummyHead := &ListNode{}
	newCur := dummyHead
	for i := 0; i < n; i++ {
		newCur.Next = &ListNode{Val: l1.Val + l2.Val}
		newCur = newCur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		newCur.Next = &ListNode{Val: l1.Val}
		newCur = newCur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		newCur.Next = &ListNode{Val: l2.Val}
		newCur = newCur.Next
		l2 = l2.Next
	}
	newCur = dummyHead.Next
	for newCur != nil {
		if newCur.Val >= 10 {
			newCur.Val %= 10
			if newCur.Next != nil {
				newCur.Next.Val += 1
			} else {
				newCur.Next = &ListNode{Val: 1}
			}
		}
		newCur = newCur.Next
	}
	return reverseList(dummyHead.Next)
}

func main() {
	l1 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	printLinkedList(addTwoNumbers(l1, l2))
	printLinkedList(addTwoNumbers(&ListNode{Val: 5}, &ListNode{Val: 5}))
}
