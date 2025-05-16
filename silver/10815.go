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

func main() {
	defer writer.Flush()

	N := scanInt()
	cards := make(map[int]bool, N)

	for i := 0; i < N; i++ {
		cards[scanInt()] = true
	}

	M := scanInt()
	for i := 0; i < M; i++ {
		if cards[scanInt()] {
			writer.WriteString("1 ")
		} else {
			writer.WriteString("0 ")
		}
	}
	writer.WriteByte('\n')
}
