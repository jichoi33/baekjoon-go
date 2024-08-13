package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

var (
	n, m    int
	numbers []int
	visited []bool
	answer  []int
)

func solve(depth int) {
	if depth == m {
		for i := 0; i < m; i++ {
			writer.WriteString(strconv.Itoa(answer[i]))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
		return
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			visited[i] = true
			answer[depth] = numbers[i]
			solve(depth + 1)
			visited[i] = false
		}
	}
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n, m = scanInt(), scanInt()

	numbers = make([]int, n)

	for i := 0; i < n; i++ {
		numbers[i] = scanInt()
	}

	sort.Ints(numbers)

	visited = make([]bool, n)
	answer = make([]int, m)

	solve(0)
}
