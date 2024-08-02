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
	graph   [][]int
	visited []bool
	n, m    int
)

func dfs(v int) {
	for _, u := range graph[v] {
		if !visited[u] {
			visited[u] = true
			dfs(u)
		}
	}
}

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n, m = nextInt(), nextInt()
	graph = make([][]int, n+1)
	visited = make([]bool, n+1)

	for i := 0; i < m; i++ {
		u, v := nextInt(), nextInt()
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	connectedComponentsCount := 0

	for i := 1; i <= n; i++ {
		if !visited[i] {
			visited[i] = true
			connectedComponentsCount++
			dfs(i)
		}
	}

	writer.WriteString(strconv.Itoa(connectedComponentsCount) + "\n")
}
