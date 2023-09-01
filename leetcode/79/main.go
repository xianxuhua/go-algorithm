package main

func inArea(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

var direct [][]int
var visited [][]bool
var m, n int

func searchWord(board [][]byte, word string, index, startX, startY int) bool {
	if index == len(word)-1 {
		return word[index] == board[startX][startY]
	}

	if word[index] == board[startX][startY] {
		visited[startX][startY] = true
		for i := 0; i < 4; i++ {
			newX := startX + direct[i][0]
			newY := startY + direct[i][1]
			if inArea(newX, newY) && !visited[newX][newY] {
				if searchWord(board, word, index+1, newX, newY) {
					return true
				}
			}
		}

		visited[startX][startY] = false
	}
	return false
}

func exist(board [][]byte, word string) bool {
	direct = [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	visited = make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	m, n = len(board), len(board[0])
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if searchWord(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func main() {

}
