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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	s := scanString()

	for i := 1; i <= len(s); i++ {
		if int(s[i-1]-'0') != i {
			writer.WriteString("-1\n")
			return
		}
	}

	writer.WriteString(strconv.Itoa(len(s)))
	writer.WriteByte('\n')
}
