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

	slogan := scanString()
	switch slogan {
	case "SONGDO":
		writer.WriteString("HIGHSCHOOL\n")
	case "CODE":
		writer.WriteString("MASTER\n")
	case "2023":
		writer.WriteString("0611\n")
	case "ALGORITHM":
		writer.WriteString("CONTEST\n")
	}
}
