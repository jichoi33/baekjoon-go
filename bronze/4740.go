package main

import (
	"bufio"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	for {
		word := scanString()
		if word == "***" {
			break
		}

		for i := len(word) - 1; i >= 0; i-- {
			writer.WriteRune(rune(word[i]))
		}
		writer.WriteByte('\n')
	}
}
