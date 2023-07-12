package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) Add(Val int) {
	pre := t.Root
	if pre == nil {
		t.Root = &Node{Val: Val}
		return
	}

	for pre.Left != nil && pre.Right != nil {
		for Val <= pre.Val && pre.Left != nil {
			pre = pre.Left
		}
		for Val >= pre.Val && pre.Right != nil {
			pre = pre.Right
		}
	}

	if Val <= pre.Val && pre.Left == nil {
		pre.Left = &Node{Val: Val}
	}
	if Val >= pre.Val && pre.Right == nil {
		pre.Right = &Node{Val: Val}
	}
}

func (t *Tree) AddWithTraverse(Val int) {
	t.Root = t.addWithTraverse(t.Root, Val)
}

// 我还是不会写递归
func (t *Tree) addWithTraverse(node *Node, Val int) *Node {
	if node == nil {
		return &Node{Val: Val}
	}

	if node.Val >= Val {
		node.Left = t.addWithTraverse(node.Left, Val)
	} else if node.Val <= Val {
		node.Right = t.addWithTraverse(node.Right, Val)
	}
	return node
}

func (t *Tree) PreTraverse(opt func(int)) {
	t.preTraverse(t.Root, opt)
}

func (t *Tree) preTraverse(now *Node, opt func(int)) {
	if now == nil {
		return
	}

	opt(now.Val)
	t.preTraverse(now.Left, opt)
	t.preTraverse(now.Right, opt)
}

func (t *Tree) PostTraverse(opt func(int)) {
	t.postTraverse(t.Root, opt)
}

func (t *Tree) postTraverse(now *Node, opt func(int)) {
	if now == nil {
		return
	}

	t.postTraverse(now.Left, opt)
	t.postTraverse(now.Right, opt)
	opt(now.Val)
}

func (t *Tree) MidTraverse(opt func(int)) {
	t.midTraverse(t.Root, opt)
}

func (t *Tree) midTraverse(now *Node, opt func(int)) {
	if now == nil {
		return
	}

	t.midTraverse(now.Left, opt)
	opt(now.Val)
	t.midTraverse(now.Right, opt)
}

func (t *Tree) midTraverseWithChannel() chan int {
	out := make(chan int)
	go func() {
		t.MidTraverse(func(i int) {
			out <- i
		})
		close(out)
	}()
	return out
}

func (t *Tree) IsBST() bool {
	Vals := []int{}
	t.MidTraverse(func(i int) {
		Vals = append(Vals, i)
	})
	for i := 0; i < len(Vals)-1; i++ {
		if Vals[i+1] < Vals[i] {
			return false
		}
	}

	return true
}

func (t *Tree) Contains(Val int) bool {
	return t.contains(t.Root, Val)
}

func (t *Tree) contains(node *Node, Val int) bool {
	if node == nil {
		return false
	}
	if node.Val == Val {
		return true
	}

	if node.Val >= Val {
		return t.contains(node.Left, Val)
	} else if node.Val <= Val {
		return t.contains(node.Right, Val)
	}
	return false
}

func main() {
	t := Tree{}
	for i := 0; i < 10; i++ {
		t.AddWithTraverse(rand.Intn(10))
	}
	t.MidTraverse(func(i int) {
		fmt.Printf("%v, ", i)
	})
	fmt.Println()
	fmt.Println(t.IsBST())
	t.PreTraverse(func(i int) {
		fmt.Printf("%v, ", i)
	})
	fmt.Println()
	t.PostTraverse(func(i int) {
		fmt.Printf("%v, ", i)
	})
	fmt.Println()
	fmt.Println(t.Contains(1))

	channel := t.midTraverseWithChannel()
	for i := range channel {
		fmt.Println(i)
	}
}
