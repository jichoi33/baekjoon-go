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

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	defer writer.Flush()

	count := 0

	N := scanInt()
	for n := 0; n < N; n++ {
		s := scanString()
		if isPalindrome(s) {
			count++
		}
	}

	if count <= 1 {
		writer.WriteString("0\n")
		return
	}

	writer.WriteString(strconv.Itoa(count * (count - 1)))
	writer.WriteByte('\n')
}
