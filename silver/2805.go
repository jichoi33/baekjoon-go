package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

var (
	n, m  int
	trees []int
)

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, m = nextInt(), nextInt()
	trees = make([]int, n)
	var maxHeight int

	for i := 0; i < n; i++ {
		trees[i] = nextInt()
		if trees[i] > maxHeight {
			maxHeight = trees[i]
		}
	}

	low, high := 0, maxHeight
	for {
		mid := (low + high + 1) / 2
		sum := 0

		for i := 0; i < n; i++ {
			if trees[i] > mid {
				sum += trees[i] - mid
			}
		}

		if low == mid {
			wr.WriteString(strconv.Itoa(mid) + "\n")
			return
		} else if sum >= m {
			low = mid
		} else {
			high = mid - 1
		}
	}
}
