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

	dp := make([]int, 12)
	dp[1], dp[2], dp[3] = 1, 2, 4
	for i := 4; i <= 11; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	t := nextInt()
	for i := 0; i < t; i++ {
		n := nextInt()
		wr.WriteString(strconv.Itoa(dp[n]) + "\n")
	}
}
