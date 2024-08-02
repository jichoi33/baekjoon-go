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

	n, x := nextInt(), nextInt()

	for i := 0; i < n; i++ {
		num := nextInt()
		if num < x {
			writer.WriteString(strconv.Itoa(num) + " ")
		}
	}
}
