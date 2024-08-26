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

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	a, x := scanInt(), scanInt()
	b, y := scanInt(), scanInt()
	t := scanInt()

	writer.WriteString(strconv.Itoa(a + maxInt(t-30, 0)*x*21))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(b + maxInt(t-45, 0)*y*21))
	writer.WriteByte('\n')
}
