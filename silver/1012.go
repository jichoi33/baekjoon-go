package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

var m, n, k int
var field [][]bool
var visited [][]bool
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func dfs(y, x int) {
	visited[y][x] = true

	for k := 0; k < 4; k++ {
		ny := y + dx[k]
		nx := x + dy[k]
		if ny >= 0 && ny < n && nx >= 0 && nx < m {
			if field[ny][nx] && !visited[ny][nx] {
				dfs(ny, nx)
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	tc := nextInt()

	for t := 0; t < tc; t++ {
		m, n, k = nextInt(), nextInt(), nextInt()

		field = make([][]bool, n)
		for i := 0; i < n; i++ {
			field[i] = make([]bool, m)
		}
		visited = make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, m)
		}

		for i := 0; i < k; i++ {
			x, y := nextInt(), nextInt()
			field[y][x] = true
		}

		var count int
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if field[i][j] && !visited[i][j] {
					count++
					dfs(i, j)
				}
			}
		}

		wr.WriteString(strconv.Itoa(count) + "\n")
	}
}
