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

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	sum := 0
	for i := 0; i < 4; i++ {
		sum += nextInt()
	}

	ans := maxInt(nextInt()*4-sum, 0)

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
