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

func scanFloat64() float64 {
	scanner.Scan()
	num, _ := strconv.ParseFloat(scanner.Text(), 64)
	return num
}

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	P := scanInt()
	for p := 0; p < P; p++ {
		N, D, A, B, F := scanInt(), scanFloat64(), scanFloat64(), scanFloat64(), scanFloat64()
		writer.WriteString(strconv.Itoa(N))
		writer.WriteByte(' ')
		writer.WriteString(strconv.FormatFloat(D/(A+B)*F, 'f', -1, 64))
		writer.WriteByte('\n')
	}
}
