package main

import (
	"bufio"
	"math/big"
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

	n := nextInt()

	ans := big.NewInt(1)
	//for i := 2; i < n+1; i++ {
	//	ans.Mul(ans, big.NewInt(int64(i)))
	//}
	ans.MulRange(1, int64(n))

	wr.WriteString(ans.String() + "\n")
}
