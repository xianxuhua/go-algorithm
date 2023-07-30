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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return _minDepth(root, 1)
}

func _minDepth(root *TreeNode, depth int) int {
	if root.Left == nil && root.Right == nil {
		return depth
	}
	left := depth
	right := depth
	if root.Left != nil {
		left = _minDepth(root.Left, depth+1)
	}
	if root.Right != nil {
		right = _minDepth(root.Right, depth+1)
	}
	return int(math.Min(float64(left), float64(right)))
}

func main() {
	t := &TreeNode{
		Val: 1,

		Right: &TreeNode{
			Val: 3,
		},
	}
	depth := minDepth(t)
	fmt.Println(depth)
}
