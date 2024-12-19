package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	defer writer.Flush()

	width, height := scanInt(), scanInt()
	area := width * height

	const yardsPerBag = 4840 * 5

	bags := (area + yardsPerBag - 1) / yardsPerBag

	writer.WriteString(strconv.Itoa(bags))
	writer.WriteByte('\n')
}
