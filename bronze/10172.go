package main

import (
	"bufio"
	"os"
)

var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	writer.WriteString("|\\_/|\n|q p|   /}\n( 0 )\"\"\"\\\n|\"^\"`    |\n||_/=\\\\__|")
}
