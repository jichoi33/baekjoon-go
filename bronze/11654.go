package main

import (
	"bufio"
	"os"
	"strconv"
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
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	r := scanString()[0]

	writer.WriteString(strconv.Itoa(int(r)))
	writer.WriteByte('\n')
}
