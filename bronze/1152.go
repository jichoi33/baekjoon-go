package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

const MaxBuf int = 1_000_001

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, MaxBuf), MaxBuf)
	writer = bufio.NewWriter(os.Stdout)
}

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	input := scanStr()
	input = strings.TrimSpace(input)
	wordCount := 0

	if input == "" {
		writer.WriteString("0\n")
		return
	}

	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			wordCount++
		}
	}
	wordCount++

	writer.WriteString(strconv.Itoa(wordCount))
	writer.WriteByte('\n')
}
