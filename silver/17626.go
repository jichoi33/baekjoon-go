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

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	dp := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		dp[i] = i
		for j := 1; j*j < i+1; j++ {
			dp[i] = minInt(dp[i], dp[i-j*j]+1)
		}
	}

	wr.WriteString(strconv.Itoa(dp[n]) + "\n")
}
