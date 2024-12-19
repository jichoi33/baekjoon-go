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

	stack := make([]int, 0, n)
	sum := 0

	for i := 0; i < n; i++ {
		num := scanInt()
		if num == 0 {
			sum -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, num)
			sum += num
		}
	}

	writer.WriteString(strconv.Itoa(sum))
	writer.WriteByte('\n')
}
