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

	n := scanInt()
	cards := make([]int, 0, 500_000)

	for i := 1; i <= n; i++ {
		cards = append(cards, i)
	}

	for len(cards) > 1 {
		cards = append(cards, cards[1])
		cards = cards[2:]
	}

	writer.WriteString(strconv.Itoa(cards[0]))
	writer.WriteByte('\n')
}
