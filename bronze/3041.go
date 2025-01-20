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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

// ==============================
// Problem Solving Logic
// ==============================

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	defer writer.Flush()

	sum := 0
	for i := 0; i < 4; i++ {
		s := scanString()
		for j := 0; j < 4; j++ {
			switch s[j] {
			case 'A':
				sum += (abs(i-0) + abs(j-0))
			case 'B':
				sum += (abs(i-0) + abs(j-1))
			case 'C':
				sum += (abs(i-0) + abs(j-2))
			case 'D':
				sum += (abs(i-0) + abs(j-3))
			case 'E':
				sum += (abs(i-1) + abs(j-0))
			case 'F':
				sum += (abs(i-1) + abs(j-1))
			case 'G':
				sum += (abs(i-1) + abs(j-2))
			case 'H':
				sum += (abs(i-1) + abs(j-3))
			case 'I':
				sum += (abs(i-2) + abs(j-0))
			case 'J':
				sum += (abs(i-2) + abs(j-1))
			case 'K':
				sum += (abs(i-2) + abs(j-2))
			case 'L':
				sum += (abs(i-2) + abs(j-3))
			case 'M':
				sum += (abs(i-3) + abs(j-0))
			case 'N':
				sum += (abs(i-3) + abs(j-1))
			case 'O':
				sum += (abs(i-3) + abs(j-2))
			}
		}
	}

	writer.WriteString(strconv.Itoa(sum))
	writer.WriteByte('\n')
}
