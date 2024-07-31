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
	paper          [][]int
	countB, countW int
)

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func isSameColor(y, x, length int) bool {
	if length == 1 {
		return true
	}

	sum := 0
	for i := y; i < y+length; i++ {
		for j := x; j < x+length; j++ {
			sum += paper[i][j]
		}
	}

	return sum == length*length || sum == 0
}

func split(y, x, length int) {
	if isSameColor(y, x, length) {
		if paper[y][x] == 1 {
			countB++
		} else {
			countW++
		}
		return
	}

	split(y, x, length/2)
	split(y+length/2, x, length/2)
	split(y, x+length/2, length/2)
	split(y+length/2, x+length/2, length/2)
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	paper = make([][]int, n)

	for i := 0; i < n; i++ {
		paper[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			paper[i][j] = nextInt()
		}
	}

	split(0, 0, n)

	wr.WriteString(strconv.Itoa(countW) + "\n")
	wr.WriteString(strconv.Itoa(countB) + "\n")
}
