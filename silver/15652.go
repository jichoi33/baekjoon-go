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
	n, m   int
	answer []int
)

func solve(at, depth int) {
	if depth == m {
		for i := 0; i < m; i++ {
			writer.WriteString(strconv.Itoa(answer[i]))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
		return
	}

	for i := at; i <= n; i++ {
		answer[depth] = i
		solve(i, depth+1)
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
	answer = make([]int, m+1)

	solve(1, 0)
}
