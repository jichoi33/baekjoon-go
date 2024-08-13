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

	T := scanInt()

	for t := 0; t < T; t++ {
		a, b := scanInt(), scanInt()
		writer.WriteString(strconv.Itoa(a + b))
		writer.WriteByte('\n')
	}
}
