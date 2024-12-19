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

	numbers := make([]int, 8)
	for i := 0; i < 8; i++ {
		numbers[i] = scanInt()
	}

	isAscending := true
	isDescending := true

	for i := 1; i < 8; i++ {
		if numbers[i] < numbers[i-1] {
			isAscending = false
			break
		}
	}

	for i := 1; i < 8; i++ {
		if numbers[i] > numbers[i-1] {
			isDescending = false
			break
		}
	}

	if isAscending {
		writer.WriteString("ascending\n")
	} else if isDescending {
		writer.WriteString("descending\n")
	} else {
		writer.WriteString("mixed\n")
	}
}
