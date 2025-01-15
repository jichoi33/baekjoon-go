package main

import (
	"bufio"
	"math"
	"os"
	"sort"
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

	n := scanInt()
	if n == 0 {
		writer.WriteString("0\n")
		return
	}

	opinions := make([]float64, n)

	for i := 0; i < n; i++ {
		opinions[i] = scanFloat64()
	}

	sort.Float64s(opinions)
	excludeCount := int(math.Round(float64(n) * 0.15))

	sum := 0.0
	for i := excludeCount; i < n-excludeCount; i++ {
		sum += opinions[i]
	}

	avg := int(math.Round(sum / float64(n-2*excludeCount)))

	writer.WriteString(strconv.Itoa(avg))
	writer.WriteByte('\n')
}
