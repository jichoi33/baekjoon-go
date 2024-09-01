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

	h, m := scanInt(), scanInt()

	m -= 45

	if m < 0 {
		h--
		m += 60

		if h < 0 {
			h += 24
		}
	}

	writer.WriteString(strconv.Itoa(h))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(m))
	writer.WriteByte('\n')
}
