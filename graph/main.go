package main

import (
	"algorithm/graph/conpoment"
	"algorithm/graph/path/bst"
	"algorithm/graph/sparse"
	"fmt"
	"os"
)

func readGraphFromFile(filename string) sparse.Graph {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	var n, m int
	fmt.Fscanf(file, "%d %d", &n, &m)
	graph := sparse.InitGraph(n, false)
	for i := 0; i < m; i++ {
		var v, w int
		fmt.Fscanf(file, "%d %d", &v, &w)
		graph.AddEdge(v, w)
	}
	return graph
}

func testConnectedComponent() {
	spareGraph1 := readGraphFromFile("/Users/xxh/projects/go/algorithm/graph/g1.txt")
	spareGraph2 := readGraphFromFile("/Users/xxh/projects/go/algorithm/graph/g2.txt")
	c1 := conpoment.InitConnectedComponent(spareGraph1)
	c2 := conpoment.InitConnectedComponent(spareGraph2)
	fmt.Println("ConnectedComponent", c1.Count())
	fmt.Println("ConnectedComponent", c2.Count())
	fmt.Println("IsConnected", c1.IsConnected(0, 1))
	fmt.Println("dfs Path", c1.Path(6))
}

func testShortestPath() {
	spareGraph1 := readGraphFromFile("/Users/xxh/projects/go/algorithm/graph/g1.txt")
	shortestPath := bst.InitShortestPath(spareGraph1)
	fmt.Println("bfs Path", shortestPath.Path(0, 6))
}

func testHeap() {

}

func main() {

}
