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

	writer.WriteString("119")
	writer.WriteByte('\n')
	writer.WriteString("jongin_go")
	writer.WriteByte('\n')
}
