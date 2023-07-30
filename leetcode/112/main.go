package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	traverse(root, func(i int) {
		fmt.Println(i)
	})
	return false
}

func traverse(root *TreeNode, opt func(int)) {
	if root == nil {
		return
	}
	opt(root.Val)
	if root.Left != nil {
		traverse(root.Left, opt)
	}
	if root.Right != nil {
		traverse(root.Right, opt)
	}
}

func main() {
	t := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(hasPathSum(t, 22))
}
