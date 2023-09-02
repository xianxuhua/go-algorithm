package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Level struct {
	Node  *TreeNode
	level int
}

func reverse(arr [][]int) [][]int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}

func levelOrderBottom(root *TreeNode) [][]int {
	q := []Level{}
	q = append(q, Level{Node: root, level: 0})
	res := [][]int{}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		if t.level == len(res) {
			res = append(res, []int{})
		}
		res[t.level] = append(res[t.level], t.Node.Val)

		if t.Node.Left != nil {
			q = append(q, Level{Node: t.Node.Left, level: t.level + 1})
		}
		if t.Node.Right != nil {
			q = append(q, Level{Node: t.Node.Right, level: t.level + 1})
		}
	}
	return reverse(res)
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
	fmt.Println(levelOrderBottom(t))
}
