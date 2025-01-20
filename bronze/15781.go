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

	N, M := scanInt(), scanInt()

	highestHelmetDefense := 0
	highestVestDefense := 0

	for i := 0; i < N; i++ {
		highestHelmetDefense = max(highestHelmetDefense, scanInt())
	}
	for i := 0; i < M; i++ {
		highestVestDefense = max(highestVestDefense, scanInt())
	}

	totalMaxDefense := highestHelmetDefense + highestVestDefense

	writer.WriteString(strconv.Itoa(totalMaxDefense))
	writer.WriteByte('\n')
}
