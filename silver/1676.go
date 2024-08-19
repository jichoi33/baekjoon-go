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

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := scanInt()

	countTwo := 0
	countFive := 0

	for i := 2; i <= n; i++ {
		for temp := i; temp%2 == 0; temp /= 2 {
			countTwo++
		}
		for temp := i; temp%5 == 0; temp /= 5 {
			countFive++
		}
	}

	ans := minInt(countTwo, countFive)

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
