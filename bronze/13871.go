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

	N, C, S := scanInt(), scanInt(), scanInt()

	curStation := 1
	count := 0

	if curStation == S {
		count++
	}

	for i := 0; i < C; i++ {
		cmd := scanInt()
		if cmd == 1 {
			curStation++
			if curStation > N {
				curStation = 1
			}
		} else {
			curStation--
			if curStation < 1 {
				curStation = N
			}
		}

		if curStation == S {
			count++
		}
	}

	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')
}
