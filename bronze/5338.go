package main

import (
	"bufio"
	"os"
)

var (
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	writer.WriteString("       _.-;;-._\n")
	writer.WriteString("'-..-'|   ||   |\n")
	writer.WriteString("'-..-'|_.-;;-._|\n")
	writer.WriteString("'-..-'|   ||   |\n")
	writer.WriteString("'-..-'|_.-''-._|\n")
}
