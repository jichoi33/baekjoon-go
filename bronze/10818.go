package main

import (
	"bufio"
	"math"
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
	min, max := math.MaxInt, math.MinInt

	for i := 0; i < n; i++ {
		num := scanInt()

		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	writer.WriteString(strconv.Itoa(min))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(max))
	writer.WriteByte('\n')
}
