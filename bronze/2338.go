package main

import (
	"bufio"
	"math/big"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	a, b := new(big.Int), new(big.Int)
	scanner.Scan()
	a.SetString(scanner.Text(), 10)
	scanner.Scan()
	b.SetString(scanner.Text(), 10)

	writer.WriteString(new(big.Int).Add(a, b).String())
	writer.WriteByte('\n')
	writer.WriteString(new(big.Int).Sub(a, b).String())
	writer.WriteByte('\n')
	writer.WriteString(new(big.Int).Mul(a, b).String())
	writer.WriteByte('\n')
}
