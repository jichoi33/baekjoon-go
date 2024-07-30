package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	dp := make([]int, 101)
	dp[1], dp[2], dp[3], dp[4], dp[5] = 1, 1, 1, 2, 2
	for i := 6; i < 101; i++ {
		dp[i] = dp[i-1] + dp[i-5]
	}

	t := nextInt()
	for i := 0; i < t; i++ {
		n := nextInt()
		wr.WriteString(strconv.Itoa(dp[n]) + "\n")
	}
}
