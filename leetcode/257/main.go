package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
		return res
	}

	leftPaths := binaryTreePaths(root.Left)
	for i := 0; i < len(leftPaths); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+leftPaths[i])
	}
	rightPaths := binaryTreePaths(root.Right)
	for i := 0; i < len(rightPaths); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+rightPaths[i])
	}
	return res
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
