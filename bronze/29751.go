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

	w, h := scanInt(), scanInt()

	ans := float64(w*h) / 2

	writer.WriteString(strconv.FormatFloat(ans, 'f', 1, 64))
	writer.WriteByte('\n')
}
