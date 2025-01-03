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
	set     map[string]bool
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

	used := make([]bool, 10001)
	for i := range numbers {
		if !used[numbers[i]] && (depth == 0 || numbers[i] >= ans[depth-1]) {
			used[numbers[i]] = true
			ans[depth] = numbers[i]
			solve(depth + 1)
		}
	}
}

func main() {
	defer writer.Flush()

	N, M = scanInt(), scanInt()

	numbers = make([]int, N)
	ans = make([]int, M)
	set = make(map[string]bool, N*N)

	for i := 0; i < N; i++ {
		numbers[i] = scanInt()
	}

	slices.Sort(numbers)

	solve(0)

	for k := range set {
		writer.WriteString(k)
		writer.WriteByte('\n')
	}
}
