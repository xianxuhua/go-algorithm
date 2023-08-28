package main

import (
	"algorithm/graph/weight/dense"
	"fmt"
	"os"
)

func readGraphFromFile(filename string) dense.Graph {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	var n, m int
	fmt.Fscanf(file, "%d %d", &n, &m)
	graph := dense.InitGraph(n, false)
	for i := 0; i < m; i++ {
		var v, w int
		var weight float32
		fmt.Fscanf(file, "%d %d %f", &v, &w, &weight)
		graph.AddEdge(v, w, weight)
	}
	return graph
}

func main() {
	graph := readGraphFromFile("/Users/xxh/projects/go/algorithm/graph/weight/g1.txt")
	graph.Show()
}
