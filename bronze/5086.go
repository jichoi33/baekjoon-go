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

	for {
		a, b := nextInt(), nextInt()

		if a == 0 || b == 0 {
			return
		}

		if a%b == 0 {
			writer.WriteString("multiple\n")
		} else if b%a == 0 {
			writer.WriteString("factor\n")
		} else {
			writer.WriteString("neither\n")
		}
	}
}
