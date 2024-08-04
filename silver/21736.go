package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

var (
	n, m    int
	campus  [][]rune
	visited [][]bool
	count   int
	dy      = []int{-1, 1, 0, 0}
	dx      = []int{0, 0, -1, 1}
)

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func dfs(y, x int) {
	visited[y][x] = true

	for i := 0; i < 4; i++ {
		curY := y + dy[i]
		curX := x + dx[i]

		if curY >= 0 && curY < n && curX >= 0 && curX < m {
			if campus[curY][curX] != 'X' && !visited[curY][curX] {
				if campus[curY][curX] == 'P' {
					count++
				}
				dfs(curY, curX)
			}
		}
	}
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n, m = nextInt(), nextInt()

	campus = make([][]rune, n)
	visited = make([][]bool, n)

	for i := 0; i < n; i++ {
		campus[i] = make([]rune, m)
		visited[i] = make([]bool, m)
	}

	var startY, startX int

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := scanner.Text()
		for j := 0; j < m; j++ {
			campus[i][j] = rune(input[j])
			if input[j] == 'I' {
				startY, startX = i, j
			}
		}
	}

	dfs(startY, startX)

	if count == 0 {
		writer.WriteString("TT\n")
		return
	}
	writer.WriteString(strconv.Itoa(count) + "\n")
}
