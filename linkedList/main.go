package main

import "fmt"

type Node struct {
	v    int
	next *Node
}

type LinkedList struct {
	head *Node
}

func Init() LinkedList {
	head := Node{v: 0}
	return LinkedList{
		head: &head,
	}
}

func (l *LinkedList) addFirst(e int) {
	node := Node{v: e}
	nextNode := l.head.next
	if nextNode != nil {
		l.head.next = &node
		node.next = nextNode
	} else {
		l.head.next = &node
	}
}

func (l *LinkedList) Print() {
	now := l.head
	for now.next != nil {
		now = now.next
		fmt.Printf("%v->", now.v)
	}
	fmt.Println()
}

func (l *LinkedList) Insert(index, v int) {
	if index < 0 {
		panic("index is invalid")
	}

	pre := l.head
	for i := 0; i < index; i++ {
		pre = pre.next
		if &pre == nil {
			panic("index is invalid")
		}
	}

	node := Node{v: v}
	next := pre.next
	fmt.Println(pre.v)
	pre.next = &node
	node.next = next
}

func (l *LinkedList) Get(index int) int {
	now := l.head.next
	for i := 0; i < index; i++ {
		now = now.next
	}
	return now.v
}

func (l *LinkedList) Delete(index int) {
	pre := l.head
	for i := 0; i < index; i++ {
		pre = pre.next
	}

	if pre.next == nil {
		panic("index is invalid")
	}

	pre.next = pre.next.next
}

// 删除给定值为val的节点
func (l *LinkedList) DeleteVal(val int) {
	pre := l.head
	for pre.next != nil {
		if pre.next.v == val {
			pre.next = pre.next.next
		} else {
			pre = pre.next
		}
	}
}

func main() {
	l := Init()
	l.addFirst(1)
	l.addFirst(2)
	l.addFirst(3)
	l.addFirst(4)
	l.addFirst(4)
	l.addFirst(5)
	l.addFirst(6)
	l.Print()
	l.Insert(6, 123)
	l.Print()
	fmt.Println(l.Get(6))
	l.Delete(0)
	l.Print()
	l.DeleteVal(4)
	l.Print()
}
