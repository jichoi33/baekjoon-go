package main

import (
	"bufio"
	"math"
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

	a, b, c := scanInt(), scanInt(), scanInt()

	writer.WriteString(strconv.Itoa(a + b - c))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(a*int(math.Pow10(len(strconv.Itoa(b)))) + b - c))
	writer.WriteByte('\n')
}
