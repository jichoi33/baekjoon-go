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

func main() {
	defer writer.Flush()

	for {
		year := scanInt()
		if year == 0 {
			return
		}

		writer.WriteString(strconv.Itoa(year))
		writer.WriteByte(' ')
		if year%4 == 0 && year >= 1896 {
			if year > 2020 {
				writer.WriteString("No city yet chosen\n")
			} else if (year >= 1914 && year <= 1918) || (year >= 1939 && year <= 1945) {
				writer.WriteString("Games cancelled\n")
			} else {
				writer.WriteString("Summer Olympics\n")
			}
		} else {
			writer.WriteString("No summer games\n")
		}
	}
}
