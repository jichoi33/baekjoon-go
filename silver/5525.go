package main

import (
	"bufio"
	"os"
	"strconv"
)

const MaxBuf int = 2_000_000

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
	scanner.Buffer(make([]byte, 0, MaxBuf), MaxBuf)
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := scanInt()
	m := scanInt()
	s := scanString()

	i, count, answer := 0, 0, 0
	for i < m-2 {
		if s[i] == 'I' && s[i+1] == 'O' && s[i+2] == 'I' {
			i += 2
			count++

			if count == n {
				answer++
				count--
			}
		} else {
			i++
			count = 0
		}
	}

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteRune('\n')
}
