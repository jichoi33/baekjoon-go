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

	writer.WriteString("I'm Sexy\n")
}
