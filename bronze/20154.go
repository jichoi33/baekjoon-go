package main

import (
	"bufio"
	"os"
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
	const MaxBuf int = 1_000_001
	scanner.Buffer(make([]byte, 0, MaxBuf), MaxBuf)
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

	S := scanString()
	alphabetStrokes := []int{
		3, 2, 1, 2, 3, 3, 3, 3, 1, 1, 3, 1, 3, 3, 1, 2, 2, 2, 1, 2, 1, 1, 2, 2, 2, 1,
	}

	sum := 0
	for _, c := range S {
		sum += alphabetStrokes[c-'A']
	}

	if sum%2 == 1 {
		writer.WriteString("I'm a winner!\n")
	} else {
		writer.WriteString("You're the winner?\n")
	}
}
