package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	N := scanInt()

	s1 := strings.Repeat("@", N*5) + "\n"
	s2 := strings.Repeat("@", N) + "\n"

	for i := 0; i < N; i++ {
		writer.WriteString(s1)
	}
	for i := 0; i < N*4; i++ {
		writer.WriteString(s2)
	}
}
