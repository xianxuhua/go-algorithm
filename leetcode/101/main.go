package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return _isSymmetric(root, root)
}

func _isSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil || right == nil && left != nil || left.Val != right.Val {
		return false
	}

	return _isSymmetric(left.Left, right.Right) && _isSymmetric(left.Right, right.Left)
}

func main() {
	p := TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3}},
	}
	fmt.Println(isSymmetric(&p))
}
