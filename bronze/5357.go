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
		scanner.Scan()
		input := scanner.Text()

		writer.WriteRune(rune(input[0]))
		for i := 1; i < len(input); i++ {
			if input[i] != input[i-1] {
				writer.WriteRune(rune(input[i]))
			}
		}
		writer.WriteByte('\n')
	}
}
