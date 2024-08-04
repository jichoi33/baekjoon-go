package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	a, b := nextInt(), nextInt()

	if a > b {
		writer.WriteString(">\n")
	} else if a < b {
		writer.WriteString("<\n")
	} else {
		writer.WriteString("==\n")
	}
}
