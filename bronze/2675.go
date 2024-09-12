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

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	T := scanInt()

	for t := 0; t < T; t++ {
		r := scanInt()
		s := scanStr()

		for i := 0; i < len(s); i++ {
			for j := 0; j < r; j++ {
				writer.WriteByte(s[i])
			}
		}
		writer.WriteByte('\n')
	}
}
