package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Level struct {
	tree  *TreeNode
	level int
}

func avg(arr []int) float64 {
	var res float64
	for i := 0; i < len(arr); i++ {
		res += float64(arr[i])
	}
	return res / float64(len(arr))
}

func averageOfLevels(root *TreeNode) []float64 {
	q := []Level{}
	q = append(q, Level{level: 0, tree: root})
	temp := [][]int{}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		if t.level == len(temp) {
			temp = append(temp, []int{})
		}
		temp[t.level] = append(temp[t.level], t.tree.Val)

		if t.tree.Left != nil {
			q = append(q, Level{level: t.level + 1, tree: t.tree.Left})
		}
		if t.tree.Right != nil {
			q = append(q, Level{level: t.level + 1, tree: t.tree.Right})
		}
	}
	res := []float64{}
	for i := 0; i < len(temp); i++ {
		res = append(res, avg(temp[i]))
	}
	return res
}

func main() {

}
