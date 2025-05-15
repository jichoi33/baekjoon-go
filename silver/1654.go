package main

import (
	"bufio"
	"os"
	"slices"
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

	K, N := scanInt(), scanInt()
	cables := make([]int, K)

	for i := 0; i < K; i++ {
		cables[i] = scanInt()
	}

	low, high := 1, slices.Max(cables)+1

	for low < high {
		mid := (low + high) / 2

		count := 0
		for _, cable := range cables {
			count += (cable / mid)
		}

		if count < N {
			high = mid
		} else {
			low = mid + 1
		}
	}

	writer.WriteString(strconv.Itoa(low - 1))
	writer.WriteByte('\n')
}
