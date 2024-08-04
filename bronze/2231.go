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

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func findGenerator(num int) int {
	start := num - len(strconv.Itoa(num))*9
	for i := start; i < num; i++ {
		tmp := i
		sum := tmp
		for tmp > 0 {
			sum += tmp % 10
			tmp /= 10
		}

		if sum == num {
			return i
		}
	}

	return 0
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := nextInt()

	writer.WriteString(strconv.Itoa(findGenerator(n)) + "\n")
}
