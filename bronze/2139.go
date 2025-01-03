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

	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	leapDays := []int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	for i := 2; i <= 12; i++ {
		days[i] += days[i-1]
		leapDays[i] += leapDays[i-1]
	}

	for {
		day, month, year := scanInt(), scanInt(), scanInt()
		if year == 0 {
			return
		}

		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			writer.WriteString(strconv.Itoa(leapDays[month-1] + day))
		} else {
			writer.WriteString(strconv.Itoa(days[month-1] + day))
		}
		writer.WriteByte('\n')
	}
}
