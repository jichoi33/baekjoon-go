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

	T := scanInt()

	for t := 0; t < T; t++ {
		h := scanInt()
		scanInt()
		n := scanInt()

		roomNumber := (n%h)*100 + (n/h + 1)
		if n%h == 0 {
			roomNumber += h * 100
			roomNumber--
		}

		writer.WriteString(strconv.Itoa(roomNumber))
		writer.WriteByte('\n')
	}
}
