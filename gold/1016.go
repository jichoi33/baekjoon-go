package main

import (
	"bufio"
	"math"
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

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	min, max := scanInt(), scanInt()

	sqrtMax := int(math.Sqrt(float64(max)))

	isPrime := make([]bool, sqrtMax+1)
	for i := 2; i <= sqrtMax; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= sqrtMax; i++ {
		if isPrime[i] {
			for j := i * i; j <= sqrtMax; j += i {
				isPrime[j] = false
			}
		}
	}

	isSquareFree := make([]bool, max-min+1)
	for i := 0; i < max-min+1; i++ {
		isSquareFree[i] = true
	}

	for i := 2; i <= sqrtMax; i++ {
		if isPrime[i] {
			square := i * i
			start := min / square
			if min%square != 0 {
				start++
			}
			firstMultiple := start * square

			for j := firstMultiple; j <= max; j += square {
				isSquareFree[j-min] = false
			}
		}
	}

	squareFreeCount := 0
	for i := range isSquareFree {
		if isSquareFree[i] {
			squareFreeCount++
		}
	}

	writer.WriteString(strconv.Itoa(squareFreeCount))
	writer.WriteByte('\n')
}
