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

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	start, end := nextInt(), nextInt()
	if start > 19 {
		start -= 24
	}

	writer.WriteString(strconv.Itoa(end - start))
	writer.WriteByte('\n')
}
