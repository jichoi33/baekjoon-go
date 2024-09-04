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

	for i := 0; i < 2; i++ {
		t, f, s, p, c := scanInt(), scanInt(), scanInt(), scanInt(), scanInt()

		writer.WriteString(strconv.Itoa(t*6 + f*3 + s*2 + p*1 + c*2))
		writer.WriteByte(' ')
	}
	writer.WriteByte('\n')
}
