package main

import (
	"fmt"
	"os"
)

func readFile(filename string) [][]int {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

func printMaze(maze [][]int) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			fmt.Printf("%3d", maze[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, // 上
	{0, -1}, // 左
	{1, 0},  // 下
	{0, 1},  // 右
}

// 点相加
func (p point) add(p2 point) point {
	return point{p.i + p2.i, p.j + p2.j}
}

// 点在迷宫中的状态
func (p point) at(maze [][]int) int {
	if p.i < 0 || p.i >= len(maze) || p.j < 0 || p.j >= len(maze[p.i]) {
		return -1
	}

	return maze[p.i][p.j]
}

func walk(maze [][]int, start, end point) [][]int {
	// 记录当前点距离起始点走过的步数
	steps := make([][]int, len(maze))
	for i := range maze {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// valid
			// maze at next is 1
			// steps at next is 0
			// next != start
			if next.at(maze) == 1 || next.at(maze) == -1 {
				continue
			}
			if next.at(steps) != 0 || next.at(maze) == -1 {
				continue
			}
			if next == start {
				continue
			}

			//
			steps[next.i][next.j] = cur.at(steps) + 1
			Q = append(Q, next)
		}
	}

	return steps
}

// 删除迷宫路径
func path(steps [][]int, start, end point) []point {
	res := []point{end}
	cur := point{end.i, end.j}

	for cur != start {
		for _, dir := range dirs {
			next := cur.add(dir)

			if next == start {
				res = append(res, start)
				return reverse(res)
			}
			if next.at(steps) == -1 || (next.at(steps) == 0) || steps[cur.i][cur.j] != steps[next.i][next.j]+1 {
				continue
			}

			res = append(res, point{next.i, next.j})
			cur = next
		}
	}

	return reverse(res)
}

func reverse(arr []point) []point {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}

func main() {
	maze := readFile("bfs/maze/maze.txt")
	printMaze(maze)
	start := point{i: 0, j: 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	steps := walk(maze, start, end)
	printMaze(steps)
	pathes := path(steps, start, end)
	fmt.Println(pathes)
}
