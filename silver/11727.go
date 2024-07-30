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

	n := nextInt()

	dp := make([]int, n+2)
	dp[1], dp[2] = 1, 3
	for i := 3; i < n+1; i++ {
		dp[i] = (dp[i-1] + dp[i-2]*2) % 10007
	}

	wr.WriteString(strconv.Itoa(dp[n]) + "\n")
}
