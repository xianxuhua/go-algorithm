package sparse

import "fmt"

type Graph struct {
	// 点、边数
	n, m int
	// 是否为有向图
	directed bool
	g        [][]int
}

func InitGraph(n int, directed bool) Graph {
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, 0)
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
	fmt.Println("**********************")
	for i := 0; i < g.n; i++ {
		for j := 0; j < len(g.g[i]); j++ {
			fmt.Print(g.g[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("**********************")
}

// AddEdge add edge of point v and w
func (g *Graph) AddEdge(v, w int) {
	if g.hasEdge(v, w) {
		return
	}

	g.g[v] = append(g.g[v], w)
	if v != w && !g.directed {
		g.g[w] = append(g.g[w], v)
	}
	g.m++
}

func (g *Graph) hasEdge(v, w int) bool {
	if v >= g.n || w >= g.n {
		panic("index out of range")
	}

	for i := 0; i < len(g.g[v]); i++ {
		if g.g[v][i] == w {
			return true
		}
	}

	return false
}

// TraverseEdge return the Adjacent nodes of node v
func (g *Graph) TraverseEdge(v int) []int {
	return g.g[v]
}

type ConnectedComponent struct {
	g       Graph
	visited []bool
	count   int
	from    []int // calc path
}

func InitConnectedComponent(g Graph) ConnectedComponent {
	from := make([]int, g.n)
	for i := 0; i < g.n; i++ {
		from[i] = -1
	}
	return ConnectedComponent{
		g:       g,
		visited: make([]bool, g.n),
		count:   0,
		from:    from,
	}
}

// Count return the number of component
func (c *ConnectedComponent) Count() int {
	count := 0
	for i := 0; i < c.g.n; i++ {
		if !c.visited[i] {
			c.dfs(i)
			count++
		}
	}
	return count
}

func (c *ConnectedComponent) dfs(v int) {
	c.visited[v] = true
	edges := c.g.TraverseEdge(v)
	for _, i := range edges {
		if !c.visited[i] {
			c.from[i] = v
			c.dfs(i)
		}
	}
}

func (c *ConnectedComponent) Path(v int) []int {
	return c.from
}
