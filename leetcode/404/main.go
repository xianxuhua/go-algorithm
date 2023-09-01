package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	q := []*TreeNode{}
	q = append(q, root)
	sum := 0
	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		if t.Left != nil && t.Left.Left == nil && t.Left.Right == nil {
			sum += t.Left.Val
		}
		if t.Left != nil {
			q = append(q, t.Left)
		}
		if t.Right != nil {
			q = append(q, t.Right)
		}
	}
	return sum
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
