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

func findMinIndex(arr []int) int {
	minIndex := 1
	for i := 2; i < len(arr); i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n, m := scanInt(), scanInt()

	rows := make([]int, n+1)
	for i := 1; i <= n; i++ {
		rows[i] = scanInt()
	}

	columns := make([]int, m+1)
	columns[1] = rows[n]
	for i := 2; i <= m; i++ {
		columns[i] = scanInt()
	}

	writer.WriteString(strconv.Itoa(findMinIndex(rows)))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(findMinIndex(columns)))
	writer.WriteByte('\n')
}
