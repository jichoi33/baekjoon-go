package main

import (
	"bufio"
	"os"
	"sort"
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

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = scanInt()
	}

	sort.Ints(numbers)

	for i := 0; i < n; i++ {
		writer.WriteString(strconv.Itoa(numbers[i]))
		writer.WriteByte('\n')
	}
}
