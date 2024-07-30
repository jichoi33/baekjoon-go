package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		if i%3 == 0 {
			dp[i] = minInt(dp[i], dp[i/3]+1)
		}
		if i%2 == 0 {
			dp[i] = minInt(dp[i], dp[i/2]+1)
		}
	}

	wr.WriteString(strconv.Itoa(dp[n]) + "\n")
}
