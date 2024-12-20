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

	myMap := make(map[int]int, N)
	for i := 0; i < N; i++ {
		myMap[scanInt()]++
	}

	M := scanInt()

	for i := 0; i < M; i++ {
		writer.WriteString(strconv.Itoa(myMap[scanInt()]))
		writer.WriteByte(' ')
	}
	writer.WriteByte('\n')
}
