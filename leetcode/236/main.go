package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root == p || root == q {
		return root
	}

	leftReturn := lowestCommonAncestor(root.Left, p, q)
	rightReturn := lowestCommonAncestor(root.Right, p, q)
	if leftReturn != nil && rightReturn != nil {
		return root
	}
	if leftReturn != nil && rightReturn == nil {
		return leftReturn
	}
	if leftReturn == nil && rightReturn != nil {
		return rightReturn
	}
	return nil
}

func main() {

}
