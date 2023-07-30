package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) add(v int) {
	cur := l
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: v}
}

func (l *ListNode) print() {
	cur := l
	for cur != nil {
		fmt.Print(cur.Val, "->")
		cur = cur.Next
	}
	fmt.Println()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur1 := list1
	cur2 := list2
	newNode := &ListNode{Val: -1}
	newCur := newNode
	for cur1 != nil && cur2 != nil {
		if cur1.Val > cur2.Val {
			newCur.Next = &ListNode{Val: cur2.Val}
			cur2 = cur2.Next
		} else {
			newCur.Next = &ListNode{Val: cur1.Val}
			cur1 = cur1.Next
		}
		newCur = newCur.Next
	}
	for cur1 != nil {
		newCur.Next = &ListNode{Val: cur1.Val}
		cur1 = cur1.Next
		newCur = newCur.Next
	}
	for cur2 != nil {
		newCur.Next = &ListNode{Val: cur2.Val}
		cur2 = cur2.Next
		newCur = newCur.Next
	}

	return newNode.Next
}

func main() {
	l := ListNode{Val: -1}
	l.add(1)
	l.add(2)
	l.add(4)
	l2 := ListNode{Val: -1}
	l2.add(1)
	l2.add(3)
	l2.add(4)
	l.Next.print()
	l2.Next.print()
	l3 := mergeTwoLists(l.Next, l2.Next)
	l3.print()
}
