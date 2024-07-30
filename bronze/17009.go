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

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	a3, a2, a1 := nextInt(), nextInt(), nextInt()
	b3, b2, b1 := nextInt(), nextInt(), nextInt()

	totalPointsA := a3*3 + a2*2 + a1*1
	totalPointsB := b3*3 + b2*2 + b1*1

	if totalPointsA > totalPointsB {
		wr.WriteString("A")
	} else if totalPointsA == totalPointsB {
		wr.WriteString("T")
	} else {
		wr.WriteString("B")
	}
	wr.WriteByte('\n')
}
