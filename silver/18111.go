package main

import (
	"bufio"
	"math"
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

	n, m, b := nextInt(), nextInt(), nextInt()
	heightCount := make([]int, 256+1)

	maxHeight := math.MinInt64
	minHeight := math.MaxInt64

	for i := 0; i < n*m; i++ {
		height := nextInt()
		heightCount[height]++

		if height > maxHeight {
			maxHeight = height
		}
		if height < minHeight {
			minHeight = height
		}
	}

	minTime := math.MaxInt64
	totalHeight := 0

	for i := maxHeight; i >= minHeight; i-- {
		curBlocks := b
		time := 0
		for j := minHeight; j <= maxHeight; j++ {
			if i > j {
				curBlocks -= (i - j) * heightCount[j]
				time += (i - j) * heightCount[j]
			} else {
				curBlocks += (j - i) * heightCount[j]
				time += (j - i) * 2 * heightCount[j]
			}
		}

		if curBlocks >= 0 && time < minTime {
			minTime = time
			totalHeight = i
		}
	}

	writer.WriteString(strconv.Itoa(minTime) + " " + strconv.Itoa(totalHeight) + "\n")
}
