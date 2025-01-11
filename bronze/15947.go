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

	switch N % 14 {
	case 1, 13:
		writer.WriteString("baby")
	case 2, 0:
		writer.WriteString("sukhwan")
	case 3, 7, 11:
		if N/14 >= 3 {
			writer.WriteString("tu+ru*")
			writer.WriteString(strconv.Itoa(2 + N/14))
		} else {
			writer.WriteString("tururu")
			for i := 0; i < N/14; i++ {
				writer.WriteString("ru")
			}
		}

	case 4, 8, 12:
		if N/14 >= 4 {
			writer.WriteString("tu+ru*")
			writer.WriteString(strconv.Itoa(1 + N/14))
		} else {
			writer.WriteString("turu")
			for i := 0; i < N/14; i++ {
				writer.WriteString("ru")
			}
		}
	case 5:
		writer.WriteString("very")
	case 6:
		writer.WriteString("cute")
	case 9:
		writer.WriteString("in")
	case 10:
		writer.WriteString("bed")
	}
	writer.WriteByte('\n')
}
