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
	defer writer.Flush()

	for scanner.Scan() {
		writer.WriteString(scanner.Text())
		writer.WriteByte('\n')
	}
}
