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

	for {
		a, b := scanInt(), scanInt()
		if a == 0 || b == 0 {
			return
		}

		writer.WriteString(strconv.Itoa(a + b))
		writer.WriteByte('\n')
	}
}
