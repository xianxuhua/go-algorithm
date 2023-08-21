package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, []int{root.Val})
		return res
	}

	leftPaths := binaryTreePaths(root.Left)
	for i := 0; i < len(leftPaths); i++ {
		leftPaths[i] = append(leftPaths[i], root.Val)
		res = append(res, leftPaths[i])
	}
	rightPaths := binaryTreePaths(root.Right)
	for i := 0; i < len(rightPaths); i++ {
		rightPaths[i] = append(rightPaths[i], root.Val)
		res = append(res, rightPaths[i])
	}
	return res
}

func sum(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}

func reverse(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	paths := binaryTreePaths(root)
	for _, path := range paths {
		if sum(path) == targetSum {
			res = append(res, reverse(path))
		}
	}
	return res
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
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(pathSum(t, 22))
}
