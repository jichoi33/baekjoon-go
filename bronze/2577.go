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

	a, b, c := scanInt(), scanInt(), scanInt()
	ans := a * b * c

	counts := make([]byte, 10)

	for ans > 0 {
		counts[ans%10]++
		ans /= 10
	}

	for i := 0; i < 10; i++ {
		writer.WriteByte(counts[i] + '0')
		writer.WriteByte('\n')
	}
}
