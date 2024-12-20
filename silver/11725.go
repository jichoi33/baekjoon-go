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

func main() {
	defer writer.Flush()

	N := scanInt()

	graph := make([][]int, N+1)

	for i := 0; i < N-1; i++ {
		u, v := scanInt(), scanInt()
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	parents := make([]int, N+1)
	visited := make([]bool, N+1)

	queue := make([]int, 0, N-1)

	queue = append(queue, 1)
	visited[1] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[cur] {
			if !visited[neighbor] {
				parents[neighbor] = cur
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	for i := 2; i <= N; i++ {
		writer.WriteString(strconv.Itoa(parents[i]))
		writer.WriteByte('\n')
	}
}
