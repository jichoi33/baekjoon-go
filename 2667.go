package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var (
	n            int
	houses       [][]int
	visited      [][]bool
	houseCounts  []int
	houseCount   int
	complexCount int
)

var (
	dy = []int{-1, 1, 0, 0}
	dx = []int{0, 0, -1, 1}
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func dfs(y, x int) {
	visited[y][x] = true

	for i := 0; i < 4; i++ {
		nextY, nextX := y+dy[i], x+dx[i]

		if nextY >= 0 && nextY < n && nextX >= 0 && nextX < n {
			if houses[nextY][nextX] == 1 && !visited[nextY][nextX] {
				houseCount++
				dfs(nextY, nextX)
			}
		}
	}
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

func main() {
	defer writer.Flush()

	n = scanInt()

	houses = make([][]int, n)
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		houses[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		line := scanString()
		for j := 0; j < n; j++ {
			houses[i][j] = int(line[j] - '0')
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			houseCount = 0
			if houses[i][j] == 1 && !visited[i][j] {
				complexCount++
				houseCount++
				dfs(i, j)
				houseCounts = append(houseCounts, houseCount)
			}
		}
	}

	sort.Ints(houseCounts)

	writer.WriteString(strconv.Itoa(complexCount))
	writer.WriteByte('\n')

	for _, hc := range houseCounts {
		writer.WriteString(strconv.Itoa(hc))
		writer.WriteByte('\n')
	}
}
