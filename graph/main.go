package main

import (
	"algorithm/graph/sparse"
	"fmt"
)

func main() {
	//denseGraph := dense.InitGraph(3, false)
	//denseGraph.AddEdge(1, 2)
	//denseGraph.Show()
	//fmt.Println(denseGraph.TraverseEdge(1))

	spareGraph := sparse.InitGraph(3, false)
	spareGraph.AddEdge(1, 0)
	spareGraph.AddEdge(1, 2)
	spareGraph.Show()
	fmt.Println(spareGraph.TraverseEdge(1))
	c := sparse.InitConnectedComponent(spareGraph)
	fmt.Println("ConnectedComponent", c.Count())
	fmt.Println("Path", c.Path(2))
}
