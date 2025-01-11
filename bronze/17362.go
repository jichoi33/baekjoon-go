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

	n := scanInt()

	if n%2 == 1 {
		if (n-1)%8 == 0 {
			writer.WriteString("1\n")
		} else if (n-5)%8 == 0 {
			writer.WriteString("5\n")
		} else {
			writer.WriteString("3\n")
		}
	} else {
		if n%4 != 0 {
			n -= 2
		}
		if (n/4)%2 == 0 {
			writer.WriteString("2\n")
		} else {
			writer.WriteString("4\n")
		}
	}
}
