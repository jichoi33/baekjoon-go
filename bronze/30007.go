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

	for i := 0; i < n; i++ {
		a, b, x := nextInt(), nextInt(), nextInt()
		writer.WriteString(strconv.Itoa(a*(x-1)+b) + "\n")
	}
}
