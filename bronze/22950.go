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
	const MaxBuf int = 1_100_000
	scanner.Buffer(make([]byte, 0, MaxBuf), MaxBuf)
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

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	N := scanInt()
	M := scanBytes()
	K := scanInt()

	start := 0
	for start < N && M[start] == '0' {
		start++
	}

	if start == N {
		writer.WriteString("YES\n")
		return
	}

	trimmedLen := N - start

	if K >= trimmedLen {
		writer.WriteString("NO\n")
		return
	}

	for i := N - K; i < N; i++ {
		if M[i] == '1' {
			writer.WriteString("NO\n")
			return
		}
	}
	writer.WriteString("YES\n")
}
