package main

import (
	"bufio"
	"os"
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

	N, K := scanInt(), scanInt()
	pre := scanInt()
	operationCount := 0

	for i := 0; i < N-1; i++ {
		cur := scanInt()
		if cur <= pre {
			if cur+K <= pre {
				writer.WriteString("-1\n")
				return
			}
			cur += K
			operationCount++
		}

		pre = cur
	}

	writer.WriteString(strconv.Itoa(operationCount))
	writer.WriteByte('\n')
}
