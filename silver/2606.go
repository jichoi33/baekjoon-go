package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

var m int
var graph = make([][]int, 101)
var visited = make([]bool, 101)
var count int

func dfs(v int) {
	if visited[v] {
		return
	}
	visited[v] = true
	count++
	for _, vertex := range graph[v] {
		dfs(vertex)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	nextInt()
	m = nextInt()
	for i := 0; i < m; i++ {
		start, end := nextInt(), nextInt()
		graph[start] = append(graph[start], end)
		//graph[end] = append(graph[end], start)
	}

	dfs(1)

	wr.WriteString(strconv.Itoa(count-1) + "\n")
}
