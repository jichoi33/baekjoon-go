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
	scanner.Split(bufio.ScanLines)
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

	if N <= 2 || N > 3 {
		writer.WriteString("Woof-meow-tweet-squeek\n")
		return
	}

	fingers := make(map[string]bool, N)

	for i := 0; i < N; i++ {
		fingers[scanString()] = true
	}

	if (fingers["1 3"] || fingers["3 1"]) && (fingers["1 4"] || fingers["4 1"]) && (fingers["3 4"] || fingers["4 3"]) {
		writer.WriteString("Wa-pa-pa-pa-pa-pa-pow!\n")
	} else {
		writer.WriteString("Woof-meow-tweet-squeek\n")
	}
}
