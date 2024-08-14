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

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := scanInt()

loop:
	for i := 0; i < n; i++ {
		d := scanInt()
		coins := make([]int, d)

		for j := 0; j < d; j++ {
			coins[j] = scanInt()
		}

		writer.WriteString("Denominations: ")
		for _, coin := range coins {
			writer.WriteString(strconv.Itoa(coin))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')

		for j := 1; j < d; j++ {
			if coins[j] < coins[j-1]*2 {
				writer.WriteString("Bad coin denominations!\n\n")

				continue loop
			}
		}

		writer.WriteString("Good coin denominations!\n\n")
	}
}
