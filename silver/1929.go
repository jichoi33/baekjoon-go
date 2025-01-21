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

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	M, N := scanInt(), scanInt()

	isPrime := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		isPrime[i] = true
	}
	for i := 2; i <= N; i++ {
		if isPrime[i] {
			for j := i * i; j <= N; j += i {
				isPrime[j] = false
			}
		}
	}

	for i := M; i <= N; i++ {
		if isPrime[i] {
			writer.WriteString(strconv.Itoa(i))
			writer.WriteByte('\n')
		}
	}
}
