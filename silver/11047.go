package main

import (
	"bufio"
	"os"
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

	n, k := nextInt(), nextInt()
	coins := make([]int, n)

	for i := 0; i < n; i++ {
		coins[i] = nextInt()
	}

	coinCount := 0
	for i := n - 1; i >= 0; i-- {
		if k >= coins[i] {
			coinCount += (k / coins[i])
			k = k % coins[i]
		}
	}

	wr.WriteString(strconv.Itoa(coinCount) + "\n")
}
