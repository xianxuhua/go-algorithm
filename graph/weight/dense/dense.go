package dense

import (
	"algorithm/graph/weight/edge"
	"fmt"
)

type Graph struct {
	// 点、边数
	n, m int
	// 是否为有向图
	directed bool
	g        [][]*edge.Edge
}

func InitGraph(n int, directed bool) Graph {
	g := make([][]*edge.Edge, n)
	for i := 0; i < n; i++ {
		g[i] = make([]*edge.Edge, n)
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
			if g.g[i][j] != nil {
				fmt.Printf("%.2f   ", g.g[i][j].W)
			} else {
				fmt.Print("nil", "   ")
			}
		}
		fmt.Println()
	}
}

// AddEdge add edge of point v and w
func (g *Graph) AddEdge(v, w int, weight float32) {
	if g.hasEdge(v, w) {
		return
	}

	g.g[v][w] = &edge.Edge{I: v, J: w, W: weight}
	if !g.directed {
		g.g[w][v] = &edge.Edge{I: w, J: v, W: weight}
	}
	g.m++
}

func (g *Graph) hasEdge(v, w int) bool {
	if v >= g.n || w >= g.n {
		panic("index out of range")
	}

	return g.g[v][w] != nil
}

func (g *Graph) TraverseEdge(v int) []int {
	res := []int{}
	for i, val := range g.g[v] {
		if val != nil {
			res = append(res, i)
		}
	}
	return res
}
