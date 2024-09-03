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

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

func main() {
	defer writer.Flush()

	n := scanInt()
	graph := make([][]byte, n)

	for i := 0; i < n; i++ {
		graph[i] = make([]byte, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			graph[i][j] = scanBytes()[0]
		}
	}

	for k := 0; k < n; k++ {
		for from := 0; from < n; from++ {
			for to := 0; to < n; to++ {
				if graph[from][k] == '1' && graph[k][to] == '1' {
					graph[from][to] = '1'
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			writer.WriteByte(graph[i][j])
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
	}
}
