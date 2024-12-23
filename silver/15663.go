package main

import (
	"bufio"
	"os"
	"slices"
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

// ==============================
// Problem Solving Logic
// ==============================

var (
	N, M    int
	numbers []int
	visited []bool
	ans     []int
)

func solve(depth int) {
	if depth == M {
		for i := 0; i < M; i++ {
			writer.WriteString(strconv.Itoa(ans[i]))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
		return
	}

	used := make([]bool, 10001)
	for i := 0; i < N; i++ {
		if !visited[i] && !used[numbers[i]] {
			visited[i] = true
			used[numbers[i]] = true
			ans[depth] = numbers[i]
			solve(depth + 1)
			visited[i] = false
		}
	}
}

func main() {
	defer writer.Flush()

	N, M = scanInt(), scanInt()
	numbers = make([]int, N)

	for i := 0; i < N; i++ {
		numbers[i] = scanInt()
	}

	visited = make([]bool, N)
	ans = make([]int, M)

	slices.Sort(numbers)

	solve(0)
}
