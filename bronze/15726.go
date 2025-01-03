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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	A, B, C := scanInt(), scanInt(), scanInt()
	ans := float64(A)

	if B >= C {
		ans *= float64(B) / float64(C)
	} else {
		ans *= float64(C) / float64(B)
	}

	writer.WriteString(strconv.Itoa(int(ans)))
	writer.WriteByte('\n')
}
