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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()

	scores := make([]int, n+1)
	for i := 1; i <= n; i++ {
		scores[i] = nextInt()
	}

	if n == 1 {
		wr.WriteString(strconv.Itoa(scores[n]) + "\n")
		return
	}

	dp := make([]int, n+1)
	dp[1] = scores[1]
	dp[2] = scores[2] + dp[1]

	for i := 3; i <= n; i++ {
		dp[i] = maxInt(dp[i-3]+scores[i-1], dp[i-2]) + scores[i]
	}

	wr.WriteString(strconv.Itoa(dp[n]) + "\n")
}
