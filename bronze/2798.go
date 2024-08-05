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

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n, m := nextInt(), nextInt()

	cards := make([]int, n)

	for i := 0; i < n; i++ {
		cards[i] = nextInt()
	}

	ans := 0

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				sum := cards[i] + cards[j] + cards[k]
				if sum > ans && sum <= m {
					ans = sum
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(ans) + "\n")
}
