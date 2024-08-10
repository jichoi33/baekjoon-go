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

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	score := scanInt()

	if score >= 90 {
		writer.WriteString("A\n")
	} else if score >= 80 {
		writer.WriteString("B\n")
	} else if score >= 70 {
		writer.WriteString("C\n")
	} else if score >= 60 {
		writer.WriteString("D\n")
	} else {
		writer.WriteString("F\n")
	}
}
