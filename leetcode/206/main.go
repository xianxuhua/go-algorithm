package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	arr := []int{}
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	newL := &ListNode{Val: arr[0]}
	newHead := newL
	for i := 1; i < len(arr); i++ {
		newL.Next = &ListNode{
			Val: arr[i],
		}
		newL = newL.Next
	}
	return newHead
}

func main() {
	l := &ListNode{Val: 1, Next: &ListNode{
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
	}}
	fmt.Println(reverseList(l))
}
