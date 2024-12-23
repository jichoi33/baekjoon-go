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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	N, K := scanInt(), scanInt()
	queue := make([]int, N)

	for i := 0; i < N; i++ {
		queue[i] = i + 1
	}

	writer.WriteString("<")
	for len(queue) > 1 {
		step := (K - 1) % len(queue)

		queue = append(queue[step:], queue[:step]...)

		writer.WriteString(strconv.Itoa(queue[0]))
		writer.WriteString(", ")

		queue = queue[1:]
	}

	writer.WriteString(strconv.Itoa(queue[0]))
	writer.WriteString(">\n")
}
