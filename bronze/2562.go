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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()

	maxNum := math.MinInt
	maxIndex := 1

	for i := 1; i <= 9; i++ {
		num := scanInt()
		if num > maxNum {
			maxNum = num
			maxIndex = i
		}
	}

	writer.WriteString(strconv.Itoa(maxNum))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(maxIndex))
	writer.WriteByte('\n')
}
