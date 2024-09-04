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

	num1, num2 := scanInt(), scanInt()

	writer.WriteString(strconv.Itoa(num1 * (num2 % 10)))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(num1 * ((num2 % 100) / 10)))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(num1 * (num2 / 100)))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(num1 * num2))
	writer.WriteByte('\n')
}
