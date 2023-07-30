package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	return _maxDepth(root, depth)
}

func _maxDepth(root *TreeNode, depth int) int {
	if root.Left == nil && root.Right == nil {
		return depth
	}
	left := depth
	right := depth
	if root.Left != nil {
		left = _maxDepth(root.Left, depth+1)
	}
	if root.Right != nil {
		right = _maxDepth(root.Right, depth+1)
	}

	return int(math.Max(float64(left), float64(right)))
}

func main() {
	t := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	depth := maxDepth(t)
	fmt.Println(depth)
}
