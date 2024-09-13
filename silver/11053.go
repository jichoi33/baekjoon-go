package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	defer writer.Flush()

	n := scanInt()
	numbers := make([]int, n)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		numbers[i] = scanInt()
	}
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if numbers[j] < numbers[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(slices.Max(dp)))
	writer.WriteByte('\n')
}
