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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	positions := make([]int, 26)
	for i := 0; i < 26; i++ {
		positions[i] = -1
	}

	S := scanString()

	for i, r := range S {
		if positions[r-'a'] == -1 {
			positions[r-'a'] = i
		}
	}

	for _, i := range positions {
		writer.WriteString(strconv.Itoa(i))
		writer.WriteByte(' ')
	}
	writer.WriteByte('\n')
}
