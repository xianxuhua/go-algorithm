package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	Level int
	Node  *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	q := []Item{}
	q = append(q, Item{Level: 0, Node: root})
	res := [][]int{}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		if t.Level == len(res) {
			res = append(res, []int{})
		}
		res[t.Level] = append(res[t.Level], t.Node.Val)

		if t.Node.Left != nil {
			q = append(q, Item{Level: t.Level + 1, Node: t.Node.Left})
		}
		if t.Node.Right != nil {
			q = append(q, Item{Level: t.Level + 1, Node: t.Node.Right})
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
	fmt.Println(levelOrder(t))
}
