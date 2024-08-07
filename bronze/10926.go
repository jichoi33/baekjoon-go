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

	scanner.Scan()
	id := scanner.Text()

	writer.WriteString(id)
	writer.WriteString("??!")
}
