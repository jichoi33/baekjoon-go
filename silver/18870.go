package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type coordinate struct {
	index, value int
}

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
	coordinates := make([]coordinate, n)

	for i := 0; i < n; i++ {
		coordinates[i] = coordinate{
			index: i,
			value: nextInt(),
		}
	}

	sort.Slice(coordinates, func(i, j int) bool {
		return coordinates[i].value < coordinates[j].value
	})

	ans := make([]int, n)

	for i, c := range coordinates {
		ans[c.index] = i

		if i > 0 {
			prevIndex := coordinates[i-1].index
			if ans[c.index]-ans[prevIndex] > 1 {
				ans[c.index] = ans[prevIndex] + 1
			}

			// If the current and previous values are the same, they should have the same answer
			if coordinates[i-1].value == c.value {
				ans[c.index] = ans[prevIndex]
			}
		}
	}

	for i := 0; i < n; i++ {
		writer.WriteString(strconv.Itoa(ans[i]) + " ")
	}
}
