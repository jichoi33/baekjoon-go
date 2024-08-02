package main

import (
	"bufio"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	scanner.Scan()
	word := scanner.Text()

	for i := 0; i < len(word); i++ {
		if word[i] >= 'a' {
			writer.WriteString(string(word[i] - 32))
		} else {
			writer.WriteString(string(word[i] + 32))
		}
	}
	writer.WriteByte('\n')
}
