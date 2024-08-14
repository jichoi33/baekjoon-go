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

	s, t, d := scanInt(), scanInt(), scanInt()

	answer := d / (s * 2) * t

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
