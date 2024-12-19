package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

	N := scanInt()

	for i := 1; i <= N; i++ {
		writer.WriteString(strings.Repeat(" ", N-i))
		writer.WriteString(strings.Repeat("*", i*2-1))
		writer.WriteByte('\n')
	}
}
