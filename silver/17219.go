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

	n, m := nextInt(), nextInt()
	addrToPw := make(map[string]string, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		addr := sc.Text()
		sc.Scan()
		addrToPw[addr] = sc.Text()
	}

	for i := 0; i < m; i++ {
		sc.Scan()
		pw := addrToPw[sc.Text()]
		wr.WriteString(pw + "\n")
	}
}
