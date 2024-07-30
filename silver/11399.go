package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func nextInt() (r int) {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	timeSlices := make([]int, n)

	for i := 0; i < n; i++ {
		timeSlices[i] = nextInt()
	}

	slices.Sort(timeSlices)

	totalTime := 0
	for i := 0; i < n; i++ {
		totalTime += timeSlices[i] * (n - i)
	}

	wr.WriteString(strconv.Itoa(totalTime) + "\n")
}
