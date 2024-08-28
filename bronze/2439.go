package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := scanInt()

	for i := 1; i <= n; i++ {
		writer.WriteString(strings.Repeat(" ", n-i))
		writer.WriteString(strings.Repeat("*", i))
		writer.WriteByte('\n')
	}
}
