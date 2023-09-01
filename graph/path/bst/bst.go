package bst

import (
	"algorithm/graph/sparse"
	"algorithm/queue"
	"algorithm/stack"
)

// ShortestPath 广度优先计算无权图的最短路径
type ShortestPath struct {
	g       sparse.Graph
	visited []bool
	from    []int
	order   []int // 从原节点到每个节点的最短距离
}

func InitShortestPath(g sparse.Graph) ShortestPath {
	from := make([]int, g.N())
	order := make([]int, g.N())
	for i := 0; i < g.N(); i++ {
		from[i] = -1
		order[i] = -1
	}
	return ShortestPath{
		g:       g,
		visited: make([]bool, g.N()),
		from:    from,
		order:   order,
	}
}

func (p *ShortestPath) Path(s int, v int) []int {
	q := queue.Queue{}
	q.Push(s)
	p.visited[s] = true
	p.order[s] = 0
	for q.Len() != 0 {
		t := q.Pop()
		edges := p.g.TraverseEdge(t)
		for _, edge := range edges {
			if !p.visited[edge] {
				p.visited[edge] = true
				p.from[edge] = t
				p.order[edge] = p.order[t] + 1
				q.Push(edge)
			}
		}
	}

	st := stack.Stack{}
	for v != -1 {
		st.Push(v)
		v = p.from[v]
	}
	res := []int{}
	for st.Len() != 0 {
		res = append(res, st.Pop())
	}
	return res
}
