package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	cur := head
	res := []int{}
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	sort.Ints(res)
	newHead := &ListNode{Val: res[0]}
	newCur := newHead
	for i := 1; i < len(res); i++ {
		newCur.Next = &ListNode{Val: res[i]}
		newCur = newCur.Next
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
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	printLinkedList(sortList(l))
}
