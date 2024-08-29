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

	h := scanInt()

	if h == 0 {
		writer.WriteString("1\n")
	} else if h == 1 {
		writer.WriteString("0\n")
	} else {
		writer.WriteString(strings.Repeat("4", h%2))
		writer.WriteString(strings.Repeat("8", h/2))
		writer.WriteByte('\n')
	}
}
