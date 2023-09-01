package dijkstra

import (
	"algorithm/graph/dense"
	"algorithm/graph/weight/edge"
)

// Dijkstra 有权图单源最短路径，不能有负权边
type Dijkstra struct {
	s       int // 源点
	visited []bool
	distTo  []float32 // 到每个目标的最短距离
	from    []*edge.Edge
	graph   dense.Graph
}

func InitDijkstra(graph dense.Graph, s int) Dijkstra {
	return Dijkstra{
		s:       s,
		visited: make([]bool, graph.N()),
		distTo:  make([]float32, graph.N()),
		from:    make([]*edge.Edge, graph.N()),
		graph:   graph,
	}
}
