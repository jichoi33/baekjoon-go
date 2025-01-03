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

func main() {
	defer writer.Flush()

	N, M := scanInt(), scanInt()
	bulbs := make([]int, N+1)

	for i := 1; i <= N; i++ {
		bulbs[i] = scanInt()
	}

	for i := 0; i < M; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()

		switch a {
		case 1:
			bulbs[b] = c
		case 2:
			for j := b; j <= c; j++ {
				if bulbs[j] == 1 {
					bulbs[j] = 0
				} else {
					bulbs[j] = 1
				}
			}
		case 3:
			for j := b; j <= c; j++ {
				bulbs[j] = 0
			}
		case 4:
			for j := b; j <= c; j++ {
				bulbs[j] = 1
			}
		}
	}

	for i := 1; i <= N; i++ {
		writer.WriteString(strconv.Itoa(bulbs[i]))
		writer.WriteByte(' ')
	}
	writer.WriteByte('\n')
}
