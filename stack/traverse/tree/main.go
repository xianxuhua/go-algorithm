package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Command struct {
	command string // go or print
	Node    *TreeNode
}

func preTraverseWithStack(root *TreeNode) []int {
	stack := []Command{}
	stack = append(stack, Command{
		command: "go",
		Node:    root,
	})
	res := []int{}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.command == "print" {
			res = append(res, top.Node.Val)
		} else { // go
			if top.Node.Right != nil {
				stack = append(stack, Command{
					command: "go",
					Node:    top.Node.Right,
				})
			}
			if top.Node.Left != nil {
				stack = append(stack, Command{
					command: "go",
					Node:    top.Node.Left,
				})
			}
			stack = append(stack, Command{
				command: "print",
				Node:    top.Node,
			})
		}
	}
	return res
}

func midTraverseWithStack(root *TreeNode) []int {
	stack := []Command{}
	stack = append(stack, Command{
		command: "go",
		Node:    root,
	})
	res := []int{}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.command == "print" {
			res = append(res, top.Node.Val)
		} else { // go
			if top.Node.Right != nil {
				stack = append(stack, Command{
					command: "go",
					Node:    top.Node.Right,
				})
			}
			stack = append(stack, Command{
				command: "print",
				Node:    top.Node,
			})
			if top.Node.Left != nil {
				stack = append(stack, Command{
					command: "go",
					Node:    top.Node.Left,
				})
			}
		}
	}
	return res
}

func postTraverseWithStack(root *TreeNode) []int {
	stack := []Command{}
	stack = append(stack, Command{
		command: "go",
		Node:    root,
	})
	res := []int{}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.command == "print" {
			res = append(res, top.Node.Val)
		} else { // go
			stack = append(stack, Command{
				command: "print",
				Node:    top.Node,
			})
			if top.Node.Right != nil {
				stack = append(stack, Command{
					command: "go",
					Node:    top.Node.Right,
				})
			}
			if top.Node.Left != nil {
				stack = append(stack, Command{
					command: "go",
					Node:    top.Node.Left,
				})
			}
		}
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
	fmt.Println(preTraverseWithStack(t))
	fmt.Println(midTraverseWithStack(t))
	fmt.Println(postTraverseWithStack(t))
}
