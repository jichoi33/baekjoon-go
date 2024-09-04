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

	lookup := map[byte]byte{
		'.':  '.',
		'O':  'O',
		'-':  '|',
		'|':  '-',
		'/':  '\\',
		'\\': '/',
		'^':  '<',
		'<':  'v',
		'v':  '>',
		'>':  '^',
	}

	n, m := scanInt(), scanInt()
	input := make([]string, n)

	for i := 0; i < n; i++ {
		input[i] = scanStr()
	}

	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			writer.WriteByte(lookup[input[j][i]])
		}
		writer.WriteByte('\n')
	}
}
