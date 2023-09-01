package main

var visited [][]bool
var direct [][]int
var m, n int

func inArea(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func dfs(grid [][]byte, startX, startY int) {
	visited[startX][startY] = true
	for i := 0; i < 4; i++ {
		newX := startX + direct[i][0]
		newY := startY + direct[i][1]
		if inArea(newX, newY) && !visited[newX][newY] && grid[newX][newY] == '1' {
			dfs(grid, newX, newY)
		}
	}
}

func numIslands(grid [][]byte) int {
	visited = make([][]bool, len(grid))
	m, n = len(grid), len(grid[0])
	direct = [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] != '0' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func main() {

}
