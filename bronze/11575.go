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
	const MaxBuf int = 1_000_001
	scanner.Buffer(make([]byte, 0, MaxBuf), MaxBuf)
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

	T := scanInt()
	for t := 0; t < T; t++ {
		a, b := scanInt(), scanInt()
		s := scanString()

		for _, c := range s {
			writer.WriteByte(byte((a*int(c-'A')+b)%26) + 'A')
		}
		writer.WriteByte('\n')
	}
}
