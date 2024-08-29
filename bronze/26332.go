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

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := scanInt()

	for i := 0; i < n; i++ {
		c, p := scanInt(), scanInt()

		totalCost := c*p - (c-1)*2

		writer.WriteString(strconv.Itoa(c))
		writer.WriteByte(' ')
		writer.WriteString(strconv.Itoa(p))
		writer.WriteByte('\n')
		writer.WriteString(strconv.Itoa(totalCost))
		writer.WriteByte('\n')
	}
}
