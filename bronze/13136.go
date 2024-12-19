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

	R, C, N := scanInt(), scanInt(), scanInt()

	cctvRows := (R + N - 1) / N
	cctvCols := (C + N - 1) / N
	totalCCTVs := cctvRows * cctvCols

	writer.WriteString(strconv.Itoa(totalCCTVs))
	writer.WriteByte('\n')
}
