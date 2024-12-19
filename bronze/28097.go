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

	N := scanInt()
	totalTime := 0

	for i := 0; i < N; i++ {
		totalTime += scanInt()
	}
	totalTime += (N - 1) * 8

	days := totalTime / 24
	hours := totalTime % 24

	writer.WriteString(strconv.Itoa(days))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(hours))
	writer.WriteByte('\n')
}
