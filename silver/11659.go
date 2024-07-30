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

	n, m := nextInt(), nextInt()

	dp := make([]int, n+1)
	dp[1] = nextInt()

	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + nextInt()
	}

	for i := 0; i < m; i++ {
		start, end := nextInt(), nextInt()
		wr.WriteString(strconv.Itoa(dp[end]-dp[start-1]) + "\n")
	}
}
