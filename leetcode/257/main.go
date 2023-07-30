package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	_binaryTreePaths(root, func(node *TreeNode) {
		fmt.Println(node.Val)
	})
	return []string{}
}

func _binaryTreePaths(root *TreeNode, opt func(node *TreeNode)) {
	if root.Left == nil && root.Right == nil {
		opt(root)
	}
	if root.Left != nil {
		_binaryTreePaths(root.Left, opt)
	}
	if root.Right != nil {
		_binaryTreePaths(root.Right, opt)
	}
}

func main() {
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(binaryTreePaths(t))
}
