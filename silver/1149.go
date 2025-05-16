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

const (
	RED   = 0
	GREEN = 1
	BLUE  = 2
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()

	N := scanInt()
	dp := make([]int, 3)

	for i := 1; i <= N; i++ {
		costR := scanInt()
		costG := scanInt()
		costB := scanInt()

		newDpR := costR + min(dp[GREEN], dp[BLUE])
		newDpG := costG + min(dp[RED], dp[BLUE])
		newDpB := costB + min(dp[RED], dp[GREEN])

		dp[RED] = newDpR
		dp[GREEN] = newDpG
		dp[BLUE] = newDpB
	}

	result := min(min(dp[RED], dp[GREEN]), dp[BLUE])

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
