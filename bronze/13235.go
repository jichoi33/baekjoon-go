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
	input := scanner.Text()

	for i := 0; i < len(input); i++ {
		if input[i] != input[len(input)-i-1] {
			writer.WriteString("false\n")
			return
		}
	}

	writer.WriteString("true\n")
}
