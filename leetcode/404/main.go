package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	_sumOfLeftLeaves(root, root, func(node *TreeNode) {
		sum += node.Val
	})
	return sum
}

func _sumOfLeftLeaves(node, parent *TreeNode, opt func(treeNode *TreeNode)) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil &&
		fmt.Sprintf("%q", node) == fmt.Sprintf("%q", parent.Left) {
		opt(node)
	}
	_sumOfLeftLeaves(node.Left, node, opt)
	_sumOfLeftLeaves(node.Right, node, opt)
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
	fmt.Println(sumOfLeftLeaves(t))
}
