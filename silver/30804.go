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

	n := nextInt()
	fruits := make([]int, n)

	for i := 0; i < n; i++ {
		fruits[i] = nextInt()
	}

	fruitToCount := make(map[int]int, 9)

	left, right, maxLen := 0, 0, 0
	for right < n {
		fruitToCount[fruits[right]]++
		right++

		for len(fruitToCount) > 2 {
			fruitToCount[fruits[left]]--

			if fruitToCount[fruits[left]] == 0 {
				delete(fruitToCount, fruits[left])
			}

			left++
		}

		if (right - left) > maxLen {
			maxLen = right - left
		}
	}

	writer.WriteString(strconv.Itoa(maxLen))
	writer.WriteByte('\n')
}
