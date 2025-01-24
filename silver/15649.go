package main

import (
	"bufio"
	"os"
	"strconv"
)

// ===========================
// I/O Setup
// ===========================
var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func init() {
	scanner.Split(bufio.ScanWords)
	// const MaxBuf int = 1_000_001
	// scanner.Buffer(make([]byte, 0, MaxBuf), MaxBuf)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

// ==============================
// Problem Solving Logic
// ==============================

var (
	N, M    int
	visited []bool
	ans     []int
)

func solve(depth int) {
	if depth == M {
		for _, v := range ans {
			writer.WriteString(strconv.Itoa(v))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
		return
	}

	for i := 1; i <= N; i++ {
		if !visited[i] {
			visited[i] = true
			ans[depth] = i
			solve(depth + 1)
			visited[i] = false
		}
	}
}

func main() {
	defer writer.Flush()

	N, M = scanInt(), scanInt()
	visited = make([]bool, N+1)
	ans = make([]int, M)

	solve(0)
}
