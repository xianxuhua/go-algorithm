package main

import (
	"fmt"
	"math/rand"
	"time"
)

type UnionFind struct {
	parent           []int
	rootElementCount []int // 效率优化，以root为根的节点的个数
}

func InitUnionFind(n int) UnionFind {
	u := UnionFind{parent: make([]int, n), rootElementCount: make([]int, n)}
	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.rootElementCount[i] = 1
	}
	return u
}

func (u *UnionFind) union(i, j int) {
	iRoot := u.find(i)
	jRoot := u.find(j)
	if iRoot == jRoot {
		return
	}

	if u.rootElementCount[iRoot] < u.rootElementCount[jRoot] {
		u.parent[iRoot] = jRoot
		u.rootElementCount[jRoot] += u.rootElementCount[iRoot]
	} else {
		u.parent[jRoot] = iRoot
		u.rootElementCount[iRoot] += u.rootElementCount[jRoot]
	}
}

func (u *UnionFind) find(i int) int {
	for u.parent[i] != i {
		i = u.parent[i]
	}
	return i
}

func (u *UnionFind) isConnected(i, j int) bool {
	return u.find(i) == u.find(j)
}

func main() {
	n := 1000000
	unionFind := InitUnionFind(n)
	startTime := time.Now()
	for i := 0; i < n-1; i++ {
		unionFind.union(rand.Intn(n), rand.Intn(n))
	}
	for i := 0; i < n; i++ {
		unionFind.isConnected(rand.Intn(n), rand.Intn(n))
	}
	fmt.Println(time.Since(startTime))
}
