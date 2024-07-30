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

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	N, m, M, T, R := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()

	if m+T > M {
		wr.WriteString("-1\n")
		return
	}

	curExercise := 0
	curPulse := m
	totalTime := 0
	for curExercise < N {
		if curPulse+T > M {
			curPulse = maxInt(curPulse-R, m)
			totalTime++
			continue
		}
		curExercise++
		curPulse += T
		totalTime++
	}

	wr.WriteString(strconv.Itoa(totalTime) + "\n")
}
