package main

import (
	"bufio"
	"os"
	"strconv"
)

type coordinate struct {
	y, x int
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n, m := scanInt(), scanInt()

	maze := make([][]bool, n)
	for i := 0; i < n; i++ {
		maze[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		s := scanString()
		for j := 0; j < m; j++ {
			maze[i][j] = s[j] == '1'
		}
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = 1e9
		}
	}

	queue := []coordinate{{0, 0}}
	dp[0][0] = 1

	dy := []int{-1, 1, 0, 0}
	dx := []int{0, 0, -1, 1}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			nextY := current.y + dy[i]
			nextX := current.x + dx[i]

			if nextY >= 0 && nextY < n && nextX >= 0 && nextX < m {
				if maze[nextY][nextX] && dp[nextY][nextX] > dp[current.y][current.x]+1 {
					dp[nextY][nextX] = dp[current.y][current.x] + 1
					queue = append(queue, coordinate{nextY, nextX})
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(dp[n-1][m-1]))
	writer.WriteByte('\n')
}
