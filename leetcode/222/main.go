package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	c := 0
	_countNodes(root, func(root *TreeNode) {
		c++
	})
	return c
}

func _countNodes(root *TreeNode, opt func(root *TreeNode)) {
	if root == nil {
		return
	}
	_countNodes(root.Left, opt)
	_countNodes(root.Right, opt)
	opt(root)
}

func main() {
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(countNodes(t))
}
