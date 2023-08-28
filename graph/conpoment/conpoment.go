package conpoment

import (
	"algorithm/graph/sparse"
	"algorithm/stack"
)

// ConnectedComponent dfs计算连通分量
type ConnectedComponent struct {
	g       sparse.Graph
	visited []bool
	count   int
	from    []int // calc path
	id      []int // 所在的连通分量
}

func InitConnectedComponent(g sparse.Graph) ConnectedComponent {
	from := make([]int, g.N())
	id := make([]int, g.N())
	for i := 0; i < g.N(); i++ {
		from[i] = -1
		id[i] = -1
	}
	return ConnectedComponent{
		g:       g,
		visited: make([]bool, g.N()),
		id:      id,
		count:   0,
		from:    from,
	}
}

// Count return the number of component
func (c *ConnectedComponent) Count() int {
	for i := 0; i < c.g.N(); i++ {
		if !c.visited[i] {
			c.dfs(i)
			c.count++
		}
	}
	return c.count
}

func (c *ConnectedComponent) IsConnected(v, w int) bool {
	return c.id[v] == c.id[w]
}

func (c *ConnectedComponent) dfs(v int) {
	c.visited[v] = true
	c.id[v] = c.count
	edges := c.g.TraverseEdge(v)
	for _, i := range edges {
		if !c.visited[i] {
			c.from[i] = v
			c.dfs(i)
		}
	}
}

// Path 查询原点到v的路径
func (c *ConnectedComponent) Path(v int) []int {
	c.dfs(v)
	s := stack.Stack{}
	p := v
	for p != -1 {
		s.Push(p)
		p = c.from[p]
	}
	res := []int{}
	for s.Len() != 0 {
		res = append(res, s.Pop())
	}
	return res
}
