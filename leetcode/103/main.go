package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	Level int
	Node  *TreeNode
}

func reverse(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}

func zigzagLevelOrder(root *TreeNode) [][]int {
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
	for i := 0; i < len(res); i++ {
		if i%2 == 1 {
			res[i] = reverse(res[i])
		}
	}
	return res
}

func main() {

}
