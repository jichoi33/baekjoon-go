package main

import (
	"bufio"
	"os"
	"strconv"
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

type coordinate struct {
	y, x int
}

func main() {
	defer writer.Flush()

	n, m := scanInt(), scanInt()
	myMap := make([][]int, n)
	visited := make([][]bool, n)

	for i := 0; i < n; i++ {
		myMap[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	target := coordinate{}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			myMap[i][j] = scanInt()
			if myMap[i][j] == 2 {
				target.y, target.x = i, j
			}
		}
	}

	dy := []int{-1, 1, 0, 0}
	dx := []int{0, 0, -1, 1}
	queue := []coordinate{target}
	visited[target.y][target.x] = true
	distance := 0

	for len(queue) > 0 {
		initialSize := len(queue)
		for i := 0; i < initialSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			myMap[cur.y][cur.x] = distance

			for j := 0; j < 4; j++ {
				tmpY := cur.y + dy[j]
				tmpX := cur.x + dx[j]
				if tmpY >= 0 && tmpY < n && tmpX >= 0 && tmpX < m {
					if myMap[tmpY][tmpX] != 0 && !visited[tmpY][tmpX] {
						queue = append(queue, coordinate{tmpY, tmpX})
						visited[tmpY][tmpX] = true
					}
				}
			}
		}
		distance++
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if myMap[i][j] == 1 && !visited[i][j] {
				myMap[i][j] = -1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			writer.WriteString(strconv.Itoa(myMap[i][j]))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
	}
}
