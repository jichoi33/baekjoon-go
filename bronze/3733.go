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

	for scanner.Scan() {
		N, _ := strconv.Atoi(scanner.Text())
		S := scanInt()

		writer.WriteString(strconv.Itoa(S / (N + 1)))
		writer.WriteByte('\n')
	}
}
