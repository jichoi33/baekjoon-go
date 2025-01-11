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

	t, p := scanFloat64(), scanFloat64()
	var batteryUsagePerMinute float64

	if p < 20 {
		batteryUsagePerMinute = (80 + (20-p)*2) / t
	} else {
		batteryUsagePerMinute = (100 - p) / t
	}

	var ans float64
	if p > 20 {
		ans = (p - 20) / batteryUsagePerMinute
		p = 20
	}
	ans += p * 2 / batteryUsagePerMinute

	writer.WriteString(strconv.FormatFloat(ans, 'f', -1, 64))
	writer.WriteByte('\n')
}
