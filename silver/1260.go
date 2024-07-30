package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

var (
	graph   [][]int
	visited []bool
)

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func dfs(v int) {
	visited[v] = true
	wr.WriteString(strconv.Itoa(v) + " ")

	for _, n := range graph[v] {
		if !visited[n] {
			dfs(n)
		}
	}
}

func bfs(v int) {
	var queue []int
	queue = append(queue, v)

	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		if !visited[now] {
			visited[now] = true
			wr.WriteString(strconv.Itoa(now) + " ")
			for _, n := range graph[now] {
				queue = append(queue, n)
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, m, v := nextInt(), nextInt(), nextInt()
	graph = make([][]int, n+1)
	visited = make([]bool, n+1)

	for i := 0; i < m; i++ {
		start, end := nextInt(), nextInt()
		graph[start] = append(graph[start], end)
		graph[end] = append(graph[end], start)
	}

	for i := 1; i < n+1; i++ {
		sort.Ints(graph[i])
	}

	dfs(v)
	wr.WriteByte('\n')
	visited = make([]bool, n+1)
	bfs(v)
	wr.WriteByte('\n')
}
