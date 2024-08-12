package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	defer writer.Flush()

	n := scanInt()
	numbers := make([]int, 10_001)

	for i := 0; i < n; i++ {
		numbers[scanInt()]++
	}

	for i := 1; i <= 10_000; i++ {
		if numbers[i] > 0 {
			writer.WriteString(strings.Repeat(fmt.Sprintf("%d\n", i), numbers[i]))
		}
	}
}
