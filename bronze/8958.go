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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	T := scanInt()
	for t := 0; t < T; t++ {
		s := scanString()

		countO := 0
		totalScore := 0

		for _, r := range s {
			if r == 'O' {
				countO++
				totalScore += countO
			} else {
				countO = 0
			}
		}

		writer.WriteString(strconv.Itoa(totalScore))
		writer.WriteByte('\n')
	}
}
