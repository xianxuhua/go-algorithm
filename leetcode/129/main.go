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

func sumNumbers(root *TreeNode) int {
	pathes := sumPath(root)
	res := 0
	for _, v := range pathes {
		atoi, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		res += atoi
	}
	return res
}

func sumPath(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}

	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
		return res
	}

	leftPaths := sumPath(root.Left)
	rightPaths := sumPath(root.Right)
	for i := 0; i < len(leftPaths); i++ {
		res = append(res, strconv.Itoa(root.Val)+leftPaths[i])
	}
	for i := 0; i < len(rightPaths); i++ {
		res = append(res, strconv.Itoa(root.Val)+rightPaths[i])
	}
	return res
}

func main() {
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	t2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	fmt.Println(sumNumbers(t))
	fmt.Println(sumNumbers(t2))
}
