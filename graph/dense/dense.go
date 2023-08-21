package dense

import "fmt"

type Graph struct {
	// 点、边数
	n, m int
	// 是否为有向图
	directed bool
	g        [][]bool
}

func InitGraph(n int, directed bool) Graph {
	g := make([][]bool, n)
	for i := 0; i < n; i++ {
		g[i] = make([]bool, n)
	}
	return Graph{
		n:        n,
		m:        0,
		directed: directed,
		g:        g,
	}
}

// N return point number
func (g *Graph) N() int {
	return g.m
}

// M return edge number
func (g *Graph) M() int {
	return g.n
}

func (g *Graph) Show() {
	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			fmt.Print(g.g[i][j], " ")
		}
		fmt.Println()
	}
}

// AddEdge add edge of point v and w
func (g *Graph) AddEdge(v, w int) {
	if g.hasEdge(v, w) {
		return
	}

	g.g[v][w] = true
	if !g.directed {
		g.g[w][v] = true
	}
	g.m++
}

func (g *Graph) hasEdge(v, w int) bool {
	if v >= g.n || w >= g.n {
		panic("index out of range")
	}

	return g.g[v][w]
}

func (g *Graph) TraverseEdge(v int) []int {
	res := []int{}
	for i, v := range g.g[v] {
		if v == true {
			res = append(res, i)
		}
	}
	return res
}
