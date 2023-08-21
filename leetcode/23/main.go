package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	channel := root.midTraverseWithChannel()
	res := 0
	for i := 0; i < k; i++ {
		res = <-channel
	}
	return res
}

func (t *TreeNode) midTraverseWithChannel() chan int {
	out := make(chan int)
	go func() {
		t.midTraverse(t, func(i int) {
			out <- i
		})
		close(out)
	}()
	return out
}

func (t *TreeNode) midTraverse(now *TreeNode, opt func(int)) {
	if now == nil {
		return
	}

	t.midTraverse(now.Left, opt)
	opt(now.Val)
	t.midTraverse(now.Right, opt)
}

func main() {
	t := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(kthSmallest(t, 1))
}
