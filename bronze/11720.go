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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := scanInt()
	s := scanString()

	sum := 0
	for i := 0; i < n; i++ {
		sum += int(s[i] - '0')
	}

	writer.WriteString(strconv.Itoa(sum))
	writer.WriteByte('\n')
}
