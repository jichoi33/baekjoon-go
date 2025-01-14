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

	A, B := scanInt(), scanInt()

	count := 1
	for B > A {
		if B%2 == 0 {
			B /= 2
			count++
		} else if B%10 == 1 {
			B /= 10
			count++
		} else {
			break
		}
	}

	if B != A {
		writer.WriteString("-1\n")
		return
	}

	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')
}
