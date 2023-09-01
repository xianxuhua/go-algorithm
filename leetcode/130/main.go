package main

var visited [][]bool
var direct [][]int
var m, n int

func inArea(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func inEdge(x, y int) bool {
	return x == 0 || x == m-1 || y == 0 || y == n-1
}

func dfs(board [][]byte, startX, startY int) {
	visited[startX][startY] = true

	for i := 0; i < 4; i++ {
		newX := startX + direct[i][0]
		newY := startY + direct[i][1]
		if inArea(newX, newY) && !visited[newX][newY] && board[newX][newY] == 'O' {
			dfs(board, newX, newY)
		}
	}
}

func solve(board [][]byte) {
	visited = make([][]bool, len(board))
	m, n = len(board), len(board[0])
	direct = [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if !visited[i][j] && board[i][j] == 'O' && !inEdge(i, j) {
				dfs(board, i, j)
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {

}
