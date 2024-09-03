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
	set := make(map[int]bool, n)

	for i := 0; i < n; i++ {
		num := scanInt()
		set[num] = true
	}

	m := scanInt()

	for i := 0; i < m; i++ {
		num := scanInt()

		if set[num] {
			writer.WriteString("1\n")
		} else {
			writer.WriteString("0\n")
		}
	}
}
