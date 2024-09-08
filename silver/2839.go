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

	n := scanInt()

	count3 := 0
	for n%5 != 0 {
		n -= 3
		count3++

		if n < 0 {
			writer.WriteString("-1\n")
			return
		}
	}
	count5 := n / 5

	writer.WriteString(strconv.Itoa(count5 + count3))
	writer.WriteByte('\n')
}
