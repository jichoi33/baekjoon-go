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

	n := nextInt()

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			writer.WriteString("*")
		}
		writer.WriteByte('\n')
	}
}
