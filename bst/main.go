package main

import (
	"algorithm/bst/queue"
	"algorithm/bst/tree"
	"fmt"
	"math/rand"
)

func LevelTraverse(t *tree.Tree, opt func(node *tree.Node)) {
	q := queue.Queue{}
	now := t.Root
	q.Push(now)

	for q.Len() != 0 {
		now = q.Pop()
		opt(now)
		if now.Left != nil {
			q.Push(now.Left)
		}
		if now.Right != nil {
			q.Push(now.Right)
		}
	}
}

func main() {
	t := tree.Tree{}
	for i := 0; i < 3; i++ {
		t.AddWithTraverse(rand.Intn(10))
	}
	t.MidTraverse(func(i int) {
		fmt.Printf("%v, ", i)
	})
	fmt.Println()
	vals := []int{}
	LevelTraverse(&t, func(node *tree.Node) {
		vals = append(vals, node.Val)
	})
	fmt.Println(vals)
}
