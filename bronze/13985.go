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
	input := scanner.Text()

	a, b, c := input[0]-'0', input[4]-'0', input[8]-'0'

	if a+b == c {
		writer.WriteString("YES\n")
	} else {
		writer.WriteString("NO\n")
	}
}
