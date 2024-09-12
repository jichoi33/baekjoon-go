package main

import (
	"bufio"
	"math/big"
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

	scanner.Scan()
	scanner.Scan()
	a, b := new(big.Int), new(big.Int)
	a.SetString(scanStr(), 10)
	b.SetString(scanStr(), 10)

	writer.WriteString(new(big.Int).Mul(a, b).String())
	writer.WriteByte('\n')
}
