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

func factorial(n int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans *= i
	}
	return ans
}

func main() {
	defer writer.Flush()

	factorials := make([]int, 6)
	factorials[0] = 1
	for i := 1; i < len(factorials); i++ {
		factorials[i] = factorials[i-1] * i
	}

	for {
		factorialNum := scanInt()
		if factorialNum == 0 {
			break
		}

		sum := 0
		count := 1
		for factorialNum > 0 {
			digit := factorialNum % 10
			factorialNum /= 10

			sum += digit * factorials[count]
			count++
		}

		writer.WriteString(strconv.Itoa(sum))
		writer.WriteByte('\n')
	}
}
