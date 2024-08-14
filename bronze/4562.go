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
		x, y := scanInt(), scanInt()

		if x >= y {
			writer.WriteString("MMM BRAINS\n")
		} else {
			writer.WriteString("NO BRAINS\n")
		}
	}
}
