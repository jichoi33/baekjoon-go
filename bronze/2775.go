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

	var apt [15][15]int

	for i := 1; i <= 14; i++ {
		apt[0][i] = i
	}

	for i := 1; i <= 14; i++ {
		for j := 1; j <= 14; j++ {
			apt[i][j] = apt[i][j-1] + apt[i-1][j]
		}
	}

	T := scanInt()

	for t := 0; t < T; t++ {
		k, n := scanInt(), scanInt()
		writer.WriteString(strconv.Itoa(apt[k][n]))
		writer.WriteByte('\n')
	}
}
