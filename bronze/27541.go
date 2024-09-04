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

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	n := scanInt()
	s := scanStr()

	if s[n-1] == 'G' {
		writer.WriteString(s[:n-1])
		writer.WriteByte('\n')
	} else {
		writer.WriteString(s)
		writer.WriteString("G\n")
	}
}
