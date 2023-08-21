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
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}
	if root.Left != nil && root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	return int(math.Min(float64(minDepth(root.Left)),
		float64(minDepth(root.Right)))) + 1
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
