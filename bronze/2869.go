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

	a, b, v := scanInt(), scanInt(), scanInt()

	days := (v-a)/(a-b) + 1
	if (v-a)%(a-b) != 0 {
		days++
	}

	writer.WriteString(strconv.Itoa(days))
	writer.WriteByte('\n')
}
