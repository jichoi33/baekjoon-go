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
	N int
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func countZones(grid [][]byte) int {
	visited := make([][]bool, N)
	for i := 0; i < N; i++ {
		visited[i] = make([]bool, N)
	}

	dy := []int{0, 0, -1, 1}
	dx := []int{-1, 1, 0, 0}

	queue := make([]coordinate, 0, N*N)
	zoneCount := 0

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if !visited[i][j] {
				visited[i][j] = true
				queue = append(queue, coordinate{i, j})
				zoneCount++
			}

			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]

				for i := 0; i < 4; i++ {
					ny := cur.y + dy[i]
					nx := cur.x + dx[i]

					if ny >= 0 && ny < N && nx >= 0 && nx < N {
						if !visited[ny][nx] && grid[ny][nx] == grid[cur.y][cur.x] {
							visited[ny][nx] = true
							queue = append(queue, coordinate{ny, nx})
						}
					}
				}
			}
		}
	}

	return zoneCount
}

func main() {
	defer writer.Flush()

	N = scanInt()

	grid := make([][]byte, N)
	for i := 0; i < N; i++ {
		grid[i] = []byte(scanString())
	}

	gridColorBlind := make([][]byte, N)
	for i := 0; i < N; i++ {
		gridColorBlind[i] = make([]byte, N)
		for j := 0; j < N; j++ {
			if grid[i][j] == 'G' {
				gridColorBlind[i][j] = 'R'
			} else {
				gridColorBlind[i][j] = grid[i][j]
			}
		}
	}

	zoneCountOriginal := countZones(grid)
	zoneCountColorBlind := countZones(gridColorBlind)

	writer.WriteString(strconv.Itoa(zoneCountOriginal))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(zoneCountColorBlind))
	writer.WriteByte('\n')
}
